package rail_fence_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecrypt(t *testing.T) {

}


func TestBatch(t *testing.T) {
	for i := 0; i <= int(EdgeDirectionRightBottom2Left); i++ {
		for j := 0; j <= int(EdgeDirectionRightBottom2Left); j++ {
			if i == j {
				continue
			}
			options := NewOptions()
			options.PutEdgeDirection = EdgeDirection(i)
			options.TakeEdgeDirection = EdgeDirection(j)

			for i := 2; i <= 10; i++ {

				options.Columns = i

				plaintext := "helloworld"
				encrypt, err := Encrypt(plaintext)
				assert.Nil(t, err)

				decrypt, err := Decrypt(encrypt)
				assert.Nil(t, err)
				assert.Equal(t, plaintext, decrypt)
			}

		}
	}
}

