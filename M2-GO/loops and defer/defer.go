
package main

import (
	"os"
	"log"
)


func createFile(fname, value string) {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	// 
	// ...
	// ...
	f.Write([]byte(value))
}

func main() {
	createFile("text.txt", "...")
}