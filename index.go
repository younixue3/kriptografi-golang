package main

import (
	"fmt"
)

func vigenereCipher(text string, key string, decrypt bool) string {
	var terenkripsi string
	kunciIndex := 0
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			if decrypt {
				char = (char-'a'-rune(key[kunciIndex]-'a')+26)%26 + 'a'
			} else {
				char = (char-'a'+rune(key[kunciIndex]-'a'))%26 + 'a'
			}
			kunciIndex = (kunciIndex + 1) % len(key)
		} else if char >= 'A' && char <= 'Z' {
			if decrypt {
				char = (char-'A'-rune(key[kunciIndex]-'a')+26)%26 + 'A'
			} else {
				char = (char-'A'+rune(key[kunciIndex]-'a'))%26 + 'A'
			}
			kunciIndex = (kunciIndex + 1) % len(key)
		}
		terenkripsi += string(char)
	}
	return terenkripsi
}

func main() {
	text := "SESABIRBRPPLWGIAMIABURSPKTCKRN"
	key := "RAHASIA"
	decrypted := vigenereCipher(text, key, true)
	fmt.Printf("Ciphertext: %s\n", "TEXT")
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Dekripsi: %s\n", decrypted)
}
