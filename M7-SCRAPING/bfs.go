
package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"log"
	"sync"
	"time"
	"strings"
	"regexp"
)



type Colors struct {
	native string
	purple string
	green string
}


var (
	COLORS = Colors{native: "\033[m", purple: "\033[95m", green: "\033[92m"}
	URL = "http://python.org"
	BASEURL = ""
	PROJECT_NAME = ""
	URLs = map[string]bool{}
	OutScopeURLs = map[string]bool{}
	EMAILs = map[string]bool{}
	SEEN = make(map[string]bool)
	PAGES = ""
	DEPTH = 1
	MAX_DEPTH = 2
	THREAD = 10
	TIMEOUT = 10
	TOKENS = make(chan struct{}, THREAD)
	MUTEX = &sync.Mutex{}
)


func setURLUniq(uri string) string {
	// Set Scheme
	uri = regexp.MustCompile(`https?://`).ReplaceAllString(uri, "http://")
	// Remove last slash!
	uri = regexp.MustCompile(`/$`).ReplaceAllString(uri, "")
	return urlCleanup(uri)
}

func urlSanitize(uri string) string {
	urparse, err := url.Parse(uri)
	if err != nil {
		uri = "http" + uri
		urparse, err = url.Parse(uri)
		if err != nil {
			return ""
		}
	}
	if urparse.Scheme == "" {
		uri = strings.Replace(uri, "://", "", -1)
		uri = "http://" + uri
	}
	return uri
}

func urlCleanup(uri string) string{
	var pos int
	// remove the spaces
	pos = strings.Index(uri, " ")
	if pos > -1 {
		uri = uri[:pos]
	}
	// remove the user@..
	pos = strings.Index(uri, "@")
	if pos > -1 {
		uri = uri[:pos]
	}
	// remove the comments
	pos = strings.Index(uri, "#")
	if pos > -1 {
		uri = uri[:pos]
	}
	return uri
}

func addEmail(uri string) bool {
	if strings.HasPrefix(strings.ToLower(uri), "mailto:") && strings.Contains(uri, "@") {
		uri = strings.ToLower(strings.ReplaceAll(uri[7:], "//", ""))
		if strings.Contains(uri, "?") {
			uri = strings.Split(uri, "?")[0]
		}
		EMAILs[uri] = true
		return true
	}
	return false
}

func urjoin(baseurl, uri string) string {
	urlower := strings.ToLower(uri)
	baseurl = strings.ReplaceAll(baseurl, `\/\/`, `//`)
	for _,chr := range []string{"", " ", "/", "#", "http://", "https://"} {
		if urlower == chr {
			return ""
		}
	}
	uri = urlCleanup(uri)
	if len(uri) == 0 {
		return ""
	}
	if len(uri) > 2 {
		if uri[:2] == "//" {
			return "http:" + uri
		}
	}
	if !strings.HasSuffix(baseurl, "/") {
		baseurl = baseurl + "/"
	}
	if strings.HasPrefix(uri, "://") {
		return ""
	}
	if strings.HasPrefix(uri, "//") {
		return baseurl + uri
	}
	if strings.HasPrefix(uri, "/") {
		return baseurl + uri[1:]
	}
	base, err := url.Parse(baseurl)
	if err != nil {
		return ""
	}
	final, err := base.Parse(uri)
	if err != nil {
		return ""
	}
	return final.String()
}

func isOutScope(host string) bool {
	host = strings.ToLower(host)
	host = strings.Replace(host, "www.", ".", 1)

	sh := strings.Split(PROJECT_NAME, ".")
	var suffix string
	if len(sh) > 1 {
		suffix = sh[len(sh)-2] + "." + sh[len(sh)-1]
	} else {
		suffix = PROJECT_NAME
	}
	if !regexp.MustCompile(`\b` + suffix).MatchString(host) {
		return true
	}
	return false
}

func validateURLs(urls []string) []string {
	newURLs := []string{}
	MUTEX.Lock()
	defer MUTEX.Unlock()
	for _, uri := range urls {
		join := urjoin(BASEURL, uri)
		if addEmail(join) || len(join) < 2 || !strings.Contains(join, "://") || !strings.Contains(join, ".") {
			continue
		}
		urparse, err := url.Parse(join)
		if err != nil {
			continue
		}
		if isOutScope(urparse.Host) {
			if !OutScopeURLs[join] {
				OutScopeURLs[join] = true
			}
			continue
		}
		newURLs = append(newURLs, join)
		if !URLs[join] {
			URLs[join] = true
		}
	}
	return newURLs
}

func getURLs(source string) []string {
	docs, err := goquery.NewDocumentFromReader(strings.NewReader(source))
	if err != nil {
		return []string{}
	}
	links := []string{}
	docs.Find("a").Each(func(i int, obj *goquery.Selection){
		link, ok := obj.Attr("href")
		if ok {
			links = append(links, link)
		}
	})
	links = validateURLs(links)
	return links
}

func request(uri string) (string, error){
	client := &http.Client{
		Timeout: time.Duration(TIMEOUT) * time.Second}
	httpTransport := &http.Transport{}
	client = &http.Client{Transport: httpTransport}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getSource(uri string) ([]string, error) {
	text, err := request(uri)
	if err != nil {
		return []string{}, err
	}
	log.Println(" > ", uri)
	PAGES += text
	allURLs := getURLs(text)

	return allURLs, nil
}

func crawl(uri string) []string {
	TOKENS <- struct{}{}
	list, err := getSource(uri)
	if err != nil {
		return []string{}
	}
	<-TOKENS
	if DEPTH == MAX_DEPTH {
		return []string{}
	}
	DEPTH += 1
	return list
}

func crawlIO() error {
	urparse, err := url.Parse(URL)
	if err != nil {
		fmt.Println("Invalid URL,", URL)
		return nil
	}

	PROJECT_NAME = strings.ToLower(strings.Replace(urparse.Host, "www.", "", -1))
	BASEURL = urlSanitize(PROJECT_NAME)
	URLs[urparse.String()] = true
	worklist := make(chan []string)
	n := 1 // number of panding sends to worklist
	// Start with the URL
	seeds, err := getSource(urparse.String())
	if err != nil {
		fmt.Println("Request:", err.Error())
		return nil
	}
	SEEN[urparse.String()] = true
	seeds = validateURLs(seeds)
	if MAX_DEPTH == 1 {
		return nil
	}

	/*
		https://google.com/index.php#footer
		https://google.com/index.php#header
		https://google.com/index.php
		or
		https://google.com/index.php/
		or
		http://google.com/index.php/

	*/

	go func(){ worklist <- seeds}()
	for ; n > 0; n-- {
		list := <-worklist
		for _, seed := range list {
			seed = setURLUniq(urjoin(BASEURL, seed))
			if !SEEN[seed] {
				SEEN[seed] = true
				n++
				go func(link string){
					worklist <- crawl(link)
				}(seed)
			}
		}
	}
	return nil
}


func main() {
	crawlIO()
	fmt.Println("URLs\n-------------")
	for k,_ := range URLs {
		fmt.Printf("\t%s%s%s\n", COLORS.green, k, COLORS.native)
	}
	fmt.Println("OutScopeURLs\n-------------")
	for k,_ := range OutScopeURLs {
		fmt.Printf("\t%s%s%s\n", COLORS.green, k, COLORS.native)
	}
	fmt.Println("EMAILs\n-------------")
	for k,_ := range EMAILs {
		fmt.Printf("\t%s%s%s\n", COLORS.green, k, COLORS.native)
	}
}