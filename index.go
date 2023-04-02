package main

import (
	"fmt"
	"strings"
)

func encryptTransposition(plaintext string) string {
	// Buat matriks kosong dengan ukuran yang sesuai
	ciphertext := ""
	rows := len(plaintext) / 5
	if len(plaintext)%5 != 0 {
		rows++
	}
	m := make([][]rune, rows)
	for i := range m {
		m[i] = make([]rune, 5)
	}

	// Isi matriks dengan karakter plaintext
	for i, r := range plaintext {
		m[i/5][i%5] = r
	}

	// Buat ciphertext dengan membaca karakter dari matriks secara vertikal
	for j := 0; j < 5; j++ {
		for i := 0; i < rows; i++ {
			if m[i][j] != 0 {
				ciphertext += string(m[i][j])
			}
		}
	}

	return ciphertext
}

func main() {
	plaintext := "IDUNIVERSITASMUHAMMADIYAHKALIMANTANTIMUR"
	ciphertext := encryptTransposition(strings.ToLower(plaintext))
	fmt.Println(ciphertext)
}
