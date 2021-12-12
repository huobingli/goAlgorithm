package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func ExampleNewCFBDecrypter() {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("939e08921a34ebc7d921c641edb55916c24cc2fa6f14e91b66c22a70c38d23e588c2aed3548cad5ab4baa63a214a")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s\n", ciphertext)
}
func ExampleNewCFBEncrypter() {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("hello golang")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	fmt.Printf("%x\n", ciphertext)
}
func main() {
	ExampleNewCFBDecrypter()
	ExampleNewCFBEncrypter()
}
