package activation_code

import (
	"crypto/ed25519"
	"crypto/rand"
)

func GenerateKeyPair() (publicKey []byte, privateKey []byte, err error) {
	//rand.Seed(time.Now().Unix())
	return ed25519.GenerateKey(rand.Reader)
}
