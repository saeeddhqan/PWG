

package main

import (
	"fmt"
	"bytes"
)


func rot13(text string) string {
	var output bytes.Buffer

	for _,c := range text {
		ord := int(c)

		if ord >= 'a' && ord <= 'z' {
			ord += 13
			if ord > 'z' {
				ord -= 26
			} else if ord < 'a' {
				ord += 26
			}
		} else if ord >= 'A' && ord <= 'Z' {
			ord += 13
			if ord > 'Z' {
				ord -= 26
			} else if ord < 'A' {
				ord += 26
			}
		}
		output.WriteString(string(ord))
	}
	return output.String()
}


func main() {
	msg := "This is a message to encode/decode with rot13 algorithm."
	encode := rot13(msg)
	fmt.Println("Encode:", encode)
	decode := rot13(encode)
	fmt.Println("Decode:", decode)
}