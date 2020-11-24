
package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
	"log"
	"fmt"
	"os"
	"strconv"
	"flag"
	"strings"
)


type Options struct {
	query 		string
	twitter		bool
	facebook	bool
	linkedin	bool
	num 		int
	limit		int
}

type Colors struct {
	native		string
	purple		string
	green		string
}


var (
	OPTIONS = Options{}
	COLORS = Colors{native: "\033[m", purple: "\033[95m", green: "\033[92m"}
	PARAMS = map[string]string{"num": "10", "start": "0", "ie": "utf-8", "oe": "utf-8", "q": "", "filter": "0"}
	URL 	= "https://google.com/search"
	ATTEMPTS = 0
	MAX_ATTEMPTS = 3
	SET_PAGE = func(page, num int) (int) {
		return (page - 1) * num
	}
	PAGES = ""
	LINKS = []string{}
)


func parseOptions() {
	if len(os.Args) < 3 {
		fmt.Printf(`Usage: %s -query [QUERY]
		-twitter, search in the Twitter (default=false)
		-facebook, search in the Facebook (default=false)
		-linkedin, search in the LinkedIn (default=false)
		-num, number of results for each page (default=10)
		-limit, number of pages that have to be crawled (default=1)
		`, os.Args[0])
		os.Exit(1)
	}
	flag.StringVar(&OPTIONS.query, "query", "", "Query string")
	flag.BoolVar(&OPTIONS.twitter, "twitter", false, "Search in the Twitter")
	flag.BoolVar(&OPTIONS.facebook, "facebook", false, "Search in the Facebook")
	flag.BoolVar(&OPTIONS.linkedin, "linkedin", false, "Search in the LinkedIn")
	flag.IntVar(&OPTIONS.num, "num", 10, "Number of results for each page (default=10)")
	flag.IntVar(&OPTIONS.limit, "limit", 1, "Number of pages that have to be crawled (default=1)")
	flag.Parse()
}


func request(uri string) (string, *http.Response, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, err
	}
	return string(body), resp, err
}


func extractLinks() {
	phrase1 := regexp.MustCompile(`<a href="([^"]*)" onmousedown`)
	phrase2 := regexp.MustCompile(`<a href="/url\?q=([^"]+)&amp;sa=U[-&;=\d\w]*`)
	for _, link := range phrase1.FindAllStringSubmatch(PAGES, -1) {
		cond1 := strings.Contains(strings.ToLower(link[1]), "https://accounts.google.com/")
		if !cond1 {
			LINKS = append(LINKS, link[1])
		}
	}
	if len(LINKS) < 1 {
		for _, link := range phrase2.FindAllStringSubmatch(PAGES, -1) {
			cond1 := strings.Contains(strings.ToLower(link[1]), "https://accounts.google.com/")
			if !cond1 {
				LINKS = append(LINKS, link[1])
			}
		}
	}
}


func showData(data map[string][]string) {
	for key, value := range data {
		fmt.Printf(" > %s%s%s\n", COLORS.green, key, COLORS.native)
		for _, atom := range value {
			fmt.Println("\t", atom)
		}
	}
} 

func results() {
	if OPTIONS.twitter {
		twitter_data := map[string][]string{"people": []string{}, "event": []string{}, "hashtag": []string{}, "status": []string{}}
		for _,link := range LINKS {
			if strings.Contains(link, "/i/events/") {
				twitter_data["event"] = append(twitter_data["event"], link)
			} else if strings.Contains(link, "/hashtag/") {
				twitter_data["hashtag"] = append(twitter_data["hashtag"], link)
			} else if strings.Contains(link, "/status/") {
				twitter_data["status"] = append(twitter_data["status"], link)
			} else {
				twitter_data["people"] = append(twitter_data["people"], link)
			}
		}
		showData(twitter_data)
	} else if OPTIONS.facebook {
		facebook_data := map[string][]string{"people": []string{}, "group": []string{}, "post": []string{}, "video": []string{}, "photo": []string{}}
		for _,link := range LINKS {
			if strings.Contains(link, "/group/") {
				facebook_data["group"] = append(facebook_data["group"], link)
			} else if regexp.MustCompile("/videos/?$").MatchString(link) {
				facebook_data["video"] = append(facebook_data["video"], link)
			} else if regexp.MustCompile(`/posts/?[^\w\d]*`).MatchString(link){
				facebook_data["post"] = append(facebook_data["post"], link)
			} else if regexp.MustCompile(`/photos/?[^\w\d]*`).MatchString(link){
				facebook_data["photo"] = append(facebook_data["photo"], link)
			} else {
				facebook_data["people"] = append(facebook_data["people"], link)
			}
		}
		showData(facebook_data)
	} else if OPTIONS.linkedin {
		linkedin_data := map[string][]string{"people": []string{}, "company": []string{}, "learning": []string{}}
		for _,link := range LINKS {
			if strings.Contains(link, "/company/") {
				linkedin_data["company"] = append(linkedin_data["company"], link)
			} else if regexp.MustCompile(`/learning/`).MatchString(link){
				linkedin_data["learning"] = append(linkedin_data["learning"], link)
			} else {
				linkedin_data["people"] = append(linkedin_data["people"], link)
			}
		}
		showData(linkedin_data)
	} else {
		showData(map[string][]string{"Pure Data": LINKS})
	}
}


func main() {
	parseOptions()
	if OPTIONS.twitter {
		PARAMS["q"] = "site:twitter.com " + OPTIONS.query
	} else if OPTIONS.facebook {
		PARAMS["q"] = "site:facebook.com " + OPTIONS.query
	} else if OPTIONS.linkedin {
		PARAMS["q"] = "site:linkedin.com " + OPTIONS.query
	} else {
		PARAMS["q"] = OPTIONS.query
	}
	PARAMS["num"] = strconv.Itoa(OPTIONS.num)
	page := 1
	for {
		params := url.Values{}
		for k,v := range PARAMS {
			params.Add(k, v)
		}
		uri := URL + "?" + params.Encode()
		log.Println("page", page)
		content, resp, err := request(uri)
		if err != nil {
			log.Printf("%s Connection error, %s %s\n", COLORS.purple, err.Error(), COLORS.native)
			ATTEMPTS++
			if MAX_ATTEMPTS == ATTEMPTS {
				break
			}
			continue
		}
		status := resp.StatusCode
		if status == 503 {
			log.Printf("%s Google CAPTCHA triggered.%s\n", COLORS.purple, COLORS.native)
			break
		}

		if status == 301 || status == 302 {
			redirect := resp.Header["location"]
			fmt.Println("redirect location:", redirect)
			content, resp, err = request(uri)
		}

		if status != 200 {
			log.Printf("[%s] %s Something went wrong.%s\n", resp.Status, COLORS.purple, COLORS.native)
			continue // missed in the videos
		}

		PAGES += content
		page++
		if page-1 >= OPTIONS.limit {
			break
		}
		PARAMS["start"] = strconv.Itoa(SET_PAGE(page, OPTIONS.num))
	}

	extractLinks()
	results()
}