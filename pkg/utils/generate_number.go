package utils

import (
	"crypto/rand"
	"io"
)

func GenerateRandomNumber(length int) string {
	// Tabel karakter yang boleh dipakai (hanya angka)
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	// Siapkan wadah untuk menampung hasil
	b := make([]byte, length)

	// Mengisi wadah dengan byte acak yang aman (crypto-safe)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}

	// Petakan byte acak tadi ke tabel angka kita
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}
