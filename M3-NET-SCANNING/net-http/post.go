
package main


import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	// "strings"
)


func main() {
	buf := url.Values{}
	buf.Add("user", "admin")
	buf.Add("pass", "pass!@#")
	// req, err := http.Post("http://localhost/login.php", "application/x-www-form-urlencoded", strings.NewReader(buf.Encode()))
	req, err := http.PostForm("http://localhost/login.php", buf)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln("Cannot convert body to string,", err)
	}
	fmt.Printf("%s\n", body)
	////////////////////////////////////////
}