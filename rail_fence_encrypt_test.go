package rail_fence_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plaintext := "helloworld"
	encrypt, err := Encrypt(plaintext, NewOptions().WithColumns(3))
	assert.Nil(t, err)
	t.Log(encrypt)
}

func TestEncryptW(t *testing.T) {

}