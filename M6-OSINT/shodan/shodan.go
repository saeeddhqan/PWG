
package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
	"os"
)


type hostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     int     `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	PostalCode   string  `json:"postal_code"`
	DMACode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}


type host struct {
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Hostnames []string     `json:"hostnames"`
	Location  hostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Domains   []string     `json:"domains"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	Port      int          `json:"port"`
	IPString  string       `json:"ip_str"`
}

type matches struct {
	Matches 	[]host 	   `json:"matches"`
}

func hostSearch(q, api string) (*matches, error) {
	res, err := http.Get(
		fmt.Sprintf("https://api.shodan.io/shodan/host/search?key=%s&query=%s", api, q),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var ret matches
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}


func main() {
	if len(os.Args) > 1 {
		log.Fatalln("Usage: API={YOUR_API_KEY} Q={QUERY}", os.Args[0])
	}
	api, q := os.Getenv("API"), os.Getenv("Q")
	data, _ := hostSearch(q, api)
	fmt.Println(data)
}

