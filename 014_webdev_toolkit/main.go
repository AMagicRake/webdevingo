package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {

	fmt.Println(hashEncrypt("This is some Data"))

}

func hashEncrypt(s string) string {
	//create a new seeded hash generator
	h := hmac.New(sha256.New, []byte("someKey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
