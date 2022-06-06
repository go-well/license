package activation_code

import (
	"crypto/ed25519"
	"crypto/rand"
)

func GenerateKeyPair() ([]byte, []byte, error) {
	return ed25519.GenerateKey(rand.Reader)
}

func GenerateKeyPairFile(private, public string) error {
	privateDer, publicDer, err := GenerateKeyPair()
	if err != nil {
		return err
	}

	err = WritePemFile(private, "ED25519 Private Key", privateDer)
	if err != nil {
		return err
	}
	return WritePemFile(public, "ED25519 Public Key", publicDer)
}
