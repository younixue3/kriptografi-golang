package main

import (
	"fmt"
	"strings"
)

// fungsi untuk membuat matriks kunci
func createMatrixKey(key string) [][]string {
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, "J", "I")

	matrix := make([][]string, 5)
	elements := make(map[string]bool)

	// membuat matriks 5x5
	for i := range matrix {
		matrix[i] = make([]string, 5)
	}

	// memasukkan huruf dari kunci ke dalam matriks
	idx := 0
	for i := range matrix {
		for j := range matrix[i] {
			if idx < len(key) {
				char := string(key[idx])
				if _, ok := elements[char]; !ok {
					matrix[i][j] = char
					elements[char] = true
					idx++
				} else {
					j--
					idx++
				}
			} else {
				for k := 'A'; k <= 'Z'; k++ {
					char := string(k)
					if char != "J" && !elements[char] {
						matrix[i][j] = char
						elements[char] = true
						break
					}
				}
			}
		}
	}

	return matrix
}

// fungsi untuk mencari posisi sebuah huruf dalam matriks
func findPosition(matrix [][]string, char string) (int, int) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1
}

// fungsi untuk melakukan enkripsi teks biasa menggunakan Playfair Cipher
func encrypt(plainText string, key string) string {
	matrix := createMatrixKey(key)

	plainText = strings.ReplaceAll(plainText, " ", "")
	plainText = strings.ToUpper(plainText)
	plainText = strings.ReplaceAll(plainText, "J", "I")
	plainText = strings.ReplaceAll(plainText, " ", "")

	var cipherText string
	for i := 0; i < len(plainText); i += 2 {
		var x1, y1, x2, y2 int
		char1 := string(plainText[i])
		char2 := ""

		if i+1 < len(plainText) {
			char2 = string(plainText[i+1])
		} else {
			char2 = "X"
		}

		x1, y1 = findPosition(matrix, char1)
		x2, y2 = findPosition(matrix, char2)

		if x1 == x2 { // jika kedua huruf pada baris yang sama
			cipherText += matrix[x1][(y1+1)%5] + matrix[x2][(y2+1)%5]
		} else if y1 == y2 { // jika kedua huruf pada kolom yang sama
			cipherText += matrix[(x1+1)%5][y1] + matrix[(x2+1)%5][y2]
		} else { // jika kedua huruf pada baris dan kolom yang berbeda
			cipherText += matrix[x1][y2] + matrix[x2][y1]
		}
	}

	return cipherText
}

// fungsi untuk melakukan dekripsi teks yang dienkripsi menggunakan Playfair Cipher
func decrypt(cipherText string, key string) string {
	matrix := createMatrixKey(key)
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	cipherText = strings.ToUpper(cipherText)

	var plainText string
	for i := 0; i < len(cipherText); i += 2 {
		var x1, y1, x2, y2 int
		char1 := string(cipherText[i])
		char2 := string(cipherText[i+1])
		fmt.Println(string(cipherText[i]))
		//if len(cipherText[1]) >= len(cipherText) {
		//	char2 = "x"
		//} else {
		//	char2 = string(cipherText[i])
		//}

		x1, y1 = findPosition(matrix, char1)
		x2, y2 = findPosition(matrix, char2)

		if x1 == x2 { // jika kedua huruf pada baris yang sama
			plainText += matrix[x1][(y1+4)%5] + matrix[x2][(y2+4)%5]
		} else if y1 == y2 { // jika kedua huruf pada kolom yang sama
			plainText += matrix[(x1+4)%5][y1] + matrix[(x2+4)%5][y2]
		} else { // jika kedua huruf pada baris dan kolom yang berbeda
			plainText += matrix[x1][y2] + matrix[x2][y1]
		}
	}

	return plainText
}

func main() {
	// contoh penggunaan
	key := "secretkey"
	plainText := "hello world "
	cipherText := encrypt(plainText, key)
	fmt.Println("Plain Text: ", plainText)
	fmt.Println("Cipher Text:", cipherText)

	decryptedText := decrypt(cipherText, key)
	fmt.Println("Decrypted Text:", decryptedText)
}
