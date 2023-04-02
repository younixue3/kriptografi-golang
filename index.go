package main

import (
	"fmt"
	"strings"
)

func rot13(input string) string {
	var output strings.Builder
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		output.WriteByte(c)
	}
	return output.String()
}

func main() {
	plaintext := "RICKO CAESAR"
	ciphertext := rot13(plaintext)
	fmt.Println(ciphertext)
}
