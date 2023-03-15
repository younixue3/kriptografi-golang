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
	text := "Ricko Caesar Aprilla Tiaka"
	key := "tuki tuki dam dam"
	encrypted := vigenereCipher(text, key)
	fmt.Printf("Plaintext: %s\n", text)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Enkripsi: %s\n", encrypted)
	fmt.Printf("======================\n")
	text = "Misalkan ini Plaintext"
	key = "india aca aca aye aye"
	encrypted = vigenereCipher(text, key)
	fmt.Printf("Plaintext: %s\n", text)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Enkripsi: %s\n", encrypted)
	fmt.Printf("======================\n")
	text = "Tubuhku kuat jiwaku sehat"
	key = "aku anak sehat"
	encrypted = vigenereCipher(text, key)
	fmt.Printf("Plaintext: %s\n", text)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Enkripsi: %s\n", encrypted)
}
