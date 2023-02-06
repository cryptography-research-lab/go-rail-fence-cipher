package rail_fence_cipher

import sentence_score "github.com/cryptography-research-lab/go-sentence-score"

func Crack(ciphertext string) (plaintext string, columns int) {
	maxScore := float64(-1)
	for i := 2; i < len(ciphertext); i++ {
		if len(ciphertext)%i != 0 {
			continue
		}
		decrypt, err := Decrypt(ciphertext)
		if err != nil {
			continue
		}
		_, score := sentence_score.CalculateScore(decrypt)
		if score > maxScore {
			maxScore = score
			columns = i
			plaintext = decrypt
		}
	}
	return
}
