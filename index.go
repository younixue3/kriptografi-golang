package main

import "fmt"

// text adalah parameter nilai yang akan di enkripsi dan shift adalah parameter untuk jumlah langkap penggeseran nilai enkripsi
func EncryptcaesarCipher(text string) string {
	result := ""
	// Perulangan Looping untuk setiap huruf
	for _, char := range text {
		// Validasi percabangan apakah nilai merupakan huruf kapital
		if char >= 'A' && char <= 'Z' {
			// Penggeseran nilai jika dibutuhkan
			shifted := (int(char-'A') + 9) % 26
			// Konversi nilai yang dihasilkan dari penggesran ke nilai karakter
			result += string(rune(shifted + 'A'))
		} else if char >= 'a' && char <= 'z' {
			// Penggeseran nilai jika dibutuhkan
			shifted := (int(char-'a') + 9) % 26
			// Konversi nilai yang dihasilkan dari penggesran ke nilai karakter
			result += string(rune(shifted + 'a'))
		} else {
			// Jika nilai tidak berupa huruf, maka kan di return default
			result += string(char)
		}
	}
	return result
}

func DecryptcaesarCipher(text string) string {
	result := ""
	// Perulangan Looping untuk setiap huruf
	for _, char := range text {
		// Check if the character is an uppercase letter
		if char >= 'A' && char <= 'Z' {
			// Apply the shift and wrap around if necessary
			shifted := (int(char-'A') - 9 + 26) % 26
			// Convert the shifted index back to a character
			result += string(rune(shifted + 'A'))
		} else if char >= 'a' && char <= 'z' {
			// Apply the shift and wrap around if necessary
			shifted := (int(char-'a') - 9 + 26) % 26
			// Convert the shifted index back to a character
			result += string(rune(shifted + 'a'))
		} else {
			// Leave the character as is if it's not a letter
			result += string(char)
		}
	}
	return result
}

func main() {
	plaintext := "RICKO CAESAR"
	encrypted := EncryptcaesarCipher(plaintext)
	decrypted := DecryptcaesarCipher(encrypted)
	fmt.Println("Plain Text:", plaintext)
	fmt.Println("Cipher Text:", encrypted)
	fmt.Println("Decrypt Text:", decrypted)
}
