package rail_fence_cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecrypt(t *testing.T) {

	//decrypt, err := Decrypt("KYsd3js2E{a2jda}")

	decrypt, err := Decrypt("ssC@sC1rct0atfvbf_ei{srtse#}",
		NewOptions().WithColumns(4).WithPutEdgeDirection(EdgeDirectionLeftTop2Bottom).WithTakeEdgeDirection(EdgeDirectionLeftTop2Right))
	// Output: ssctf{ssCtf_seC10ver#@rabit}

	assert.Nil(t, err)
	t.Log(decrypt)
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

func TestDecryptW(t *testing.T) {

	ciphertext := "ccehgyaefnpeoobe{lcirg}epriec_ora_g"
	// Output: cyberpeace{railfence_cipher_gogogo}

	//ciphertext, err := EncryptW("helloworld")
	//assert.Nil(t, err)
	//t.Log(ciphertext)
	decrypt, err := DecryptW(ciphertext, NewOptions().WithRows(5))
	assert.Nil(t, err)
	t.Log(decrypt)
}
