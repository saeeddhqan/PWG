
package main

import (
	"io"
	"os"
)

func main() {
	firstFile, err := os.Open("red.png")
	if err != nil {
		panic(err)
	}
	defer firstFile.Close()
	secondFile, err := os.Open("private_message.txt")
	if err != nil {
		panic(err)
	}
	defer secondFile.Close()
	stegFile, err := os.Create("steg_file.png")
	if err != nil {
		panic(err)
	}
	defer stegFile.Close()
	_, err = io.Copy(stegFile, firstFile)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(stegFile, secondFile)
	if err != nil {
		panic(err)
	}
	stegFile.Close()
}