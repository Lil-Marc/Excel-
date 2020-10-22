package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	a := "191&13298125&148dc1e0c110d03e7995d2e12234d73f"
	b := Cal32BitMD5Str(a)
	fmt.Println("md5加密后为:\n"+b)
}

func Cal32BitMD5Str(content string) (string) {
	md5 := md5.New()
	md5.Write([]byte(content))
	cipherStr := md5.Sum(nil)
	hexCipherStr := hex.EncodeToString(cipherStr)

	return hexCipherStr
}