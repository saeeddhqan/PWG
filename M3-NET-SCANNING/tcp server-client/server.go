
package main


import (
	"log"
	"bufio"
	"net"
	"strconv"
)


var (
	PORT = 2020
	ADDRESS = "localhost"
	PROTOCOL = "tcp"
)

func echo(conn net.Conn) {
	defer conn.Close()
	var code string = `HTTP/1.1 200 OK
	Accept-Ranges: bytes
	Content-Type: text/html


	<html>
		<form method=POST>
			<input type=text name=user>
			<input type=password name=pass>
			<input type=submit name=sub>
		</form>
	</html`
	buf := make([]byte, 4096)
	// Create a new reader to reseive data from the connection.
	reader := bufio.NewReader(conn)
	content, err := reader.Read(buf)
	if err != nil {
		log.Println("Unable to read data!")
		return
	}
	log.Println("Reseived 4096(max) bytes from the client:", string(buf[:content]))

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(code); err != nil {
		log.Println("Unable to write data!")
		return
	}
	writer.Flush()
}


func main() {
	server, err := net.Listen(PROTOCOL, ADDRESS + ":" + strconv.Itoa(PORT))
	if err != nil {
		log.Fatalln("Cannot bind the port,", err)
	}
	log.Printf("Server binded on %s:%d\n", ADDRESS, PORT)
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Cannot accept the request,", conn)
			continue
		}
		log.Println("Request reseived.")
		go echo(conn)
	}
}