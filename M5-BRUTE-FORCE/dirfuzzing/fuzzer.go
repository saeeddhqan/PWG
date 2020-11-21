
package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"strconv"
)


type Options struct{
	url		string
	list	string
	excludeStatusCode	string
	excludeURL			string
	verbose				bool
}

type Colors struct{
	native	string
	purple	string
	green	string
}


var OPTIONS = Options{}
var COLORS = Colors{native: "\033[m", purple: "\033[95m", green: "\033[92m"}
var doneThread = make(chan bool)
var activeThread = 0
var maxThread = 30


func parseOptions() {
	if len(os.Args) < 5 {
		fmt.Printf(`Usage: %s -url <URL> -list <LIST>
		-list, path list should be separated with newline
		-ex-code, exclude http status code. it could be one or more than one code, example, "404,403,400"
		-ex-url, exclude urls which match with -ex-url regex, example, "\?[\w\d]+=.*" which means that all of urls that contain a get parameter
		`, os.Args[0])
		os.Exit(1)
	}

	flag.StringVar(&OPTIONS.url, "url", "", "URL string")
	flag.StringVar(&OPTIONS.list, "list", "", "Username and password list")
	flag.StringVar(&OPTIONS.excludeStatusCode, "ex-code", `4\d\d|5\d\d`, "Status Code Exclude (default, all of status codes that are not in range 400 or 500")
	flag.StringVar(&OPTIONS.excludeURL, "ex-url", ".*?", "URL Exclude (default, all)")
	flag.BoolVar(&OPTIONS.verbose, "verbose", false, "Verbosity")
	flag.Parse()
}

func statusCodeExcluding(code int) bool {
	compile := regexp.MustCompile(OPTIONS.excludeStatusCode)
	if compile.MatchString(strconv.Itoa(code)) {
		return true
	}
	return false
}


func urlExcluding(uri string) bool {
	compile := regexp.MustCompile(OPTIONS.excludeURL)
	if compile.MatchString(uri) {
		return true
	}
	return false
}


func urlJoin(uri, urj string) (string, error) {
	urparse, err := url.Parse(uri)
	if err != nil {
		return uri, err
	}
	rel, err := urparse.Parse(urj)
	if err != nil {
		return uri, err
	}
	return rel.String(), err
}
 

func auth(path string, doneChan chan bool) {
	defer func(){doneChan<-true}()

	urjoin, err := urlJoin(OPTIONS.url, path)
	if err != nil {
		return 
	}

	if !urlExcluding(urjoin) {
		return
	}

	client := http.Client{}
	req, err := http.NewRequest("GET", urjoin, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err == nil {
		status := resp.StatusCode
		if !statusCodeExcluding(status) {
			log.Printf("%s %s => %s %s\n", COLORS.green, urjoin, resp.Status, COLORS.native)
		} else {
			if OPTIONS.verbose {
				log.Printf("%s %s => %s %s\n", COLORS.purple, urjoin, resp.Status, COLORS.native)
			}
		}
	}
}


func main() {
	parseOptions()
	buffer := make([]byte, 500000) // 500K(almost)
	file, err := os.Open(OPTIONS.list)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	EOB, err := file.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	list := strings.Split(string(buffer[:EOB]), "\n")
	for _,payload := range list {
		go auth(payload, doneThread)
		activeThread++
		if activeThread >= maxThread {
			<-doneThread
			activeThread -= 1
		}
	}

	for activeThread > 0 {
		<-doneThread
		activeThread -= 1
	}
}