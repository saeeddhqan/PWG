
package main

import (
	"net/http"
	"fmt"
	"regexp"
	"io/ioutil"
	"os"
)


func main() {
	// https://rest.db.ripe.net/search.json?query-string=86.57.64.115&flags=no-referenced&flags=no-irt&source=RIPE
	if len(os.Args) != 2 {
		fmt.Println("IP-Lookup:", os.Args[0], "IP|HOSTNAME")
		os.Exit(1)
	}

	name := os.Args[1]
	url := fmt.Sprintf("https://rest.db.ripe.net/search.json?query-string=%s&flags=no-referenced&flags=no-irt&source=RIPE", name)
	req, err := http.Get(url)
	if err != nil {
		fmt.Println("Connection Error,", err.Error())
		os.Exit(1)
	}
	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)
	toString := string(body)

	sanitize := regexp.MustCompile(`[\{\}"\[\],]+`).ReplaceAllString(toString, "")
	fmt.Println(sanitize)
}