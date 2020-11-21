
package main


import (
	"os"
	"net"
	"fmt"
)


func main(){
	if len(os.Args) != 2 {
		fmt.Printf("MX-Record: %s IP|HOSTNAME\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Println("Finding MX Records for ", name)
	mxRecords, err := net.LookupMX(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _,mxRecord := range mxRecords {
		fmt.Println("Host:", mxRecord.Host, "\tPrecedence:", mxRecord.Pref)
	}


}