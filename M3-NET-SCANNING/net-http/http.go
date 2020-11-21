
package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)


func main() {
	req, err := http.Get("http://localhost/login.php?user=admin&pass=pass!@%23")
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()
	// fmt.Println(req.Header)
	// for k,v := range req.Header {
	// 	fmt.Println(k, ":")
	// 	fmt.Println("\t\t", v)
	// }
	////////////////////////////////////////
	// fmt.Println(req.Status)
	////////////////////////////////////////
	// fmt.Println(req.StatusCode)
	// if req.StatusCode == 404 {
	// 	fmt.Println("Page not found!")
	// }
	////////////////////////////////////////
	// fmt.Println(req.Proto)
	////////////////////////////////////////
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln("Cannot convert body to string,", err)
	}
	fmt.Printf("%s\n", body)
	////////////////////////////////////////
}