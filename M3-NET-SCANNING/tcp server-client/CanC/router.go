
package main

import (
	"net"
	"bufio"
	"os"
	"log"
	"strconv"
)

var (
	PORT = 2020
	ADDRESS = "127.0.0.1"
	PROTOCOL = "tcp"
)

func echo(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 131072)
		read, err := conn.Read(buf)
		if err != nil {
			log.Println("Unable to read data!")
			continue
		}
		log.Println("\n", string(buf[:read]))
		log.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(command[:len(command)-1]); err != nil {
			log.Println("Unable to write data!")
			continue
		}
		writer.Flush()
	}
}


func main() {
	server, err := net.Listen(PROTOCOL, ADDRESS + ":" + strconv.Itoa(PORT))
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()
	log.Println("Server binded!")
	counter := 1
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("\nClient-%d connected.\n\n", counter)
		counter++
		go echo(conn)
	}
}