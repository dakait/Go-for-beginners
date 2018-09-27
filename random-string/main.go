package main 

import (
	"crypto/rand"
	"fmt"
)	

func generateRandomSecret(size int) string {
	alphanum := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefbgijklmnopqrstuvwxyz0123456789"
	var bytes = make([]byte, size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func main() {
	fmt.Println(generateRandomSecret(6))
}