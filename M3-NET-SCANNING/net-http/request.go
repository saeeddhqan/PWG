

package main


import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)


func main() {
	req, err := http.NewRequest("GET", "http://localhost/login.php", nil)
	if err != nil {
		log.Fatalln(req)
	}
	var client http.Client = http.Client{}
	req.Header.Set("User-Agent", "Request framework(example) V1.0.0")
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Cannot convert body to string,", err)
	}

	fmt.Printf("%s\n", body)
}