package rail_fence_cipher

import (
	"fmt"
	"testing"
)

func TestCrack(t *testing.T) {
	ciphertext := "hlodeorxlwlx"
	plaintext, columns := Crack(ciphertext)
	fmt.Println(columns)
	fmt.Println(plaintext)
}
