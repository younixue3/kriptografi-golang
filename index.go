package main

import (
	"fmt"
)

func vigenereCipher(text string, key string) string {
	var terenkripsi string
	kunciIndex := 0
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			char = (char-'a'+rune(key[kunciIndex]-'a'))%26 + 'a'
			kunciIndex = (kunciIndex + 1) % len(key)
		} else if char >= 'A' && char <= 'Z' {
			char = (char-'A'+rune(key[kunciIndex]-'a'))%26 + 'A'
			kunciIndex = (kunciIndex + 1) % len(key)
		}
		terenkripsi += string(char)
	}
	return terenkripsi
}

func main() {
	text := "SELAMAT MENJALANKAN IBADAH PUASA"
	key := "TIAKA"
	encrypted := vigenereCipher(text, key)
	fmt.Printf("Plaintext: %s\n", text)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Enkripsi: %s\n", encrypted)
}
