package activation_code

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
)

func Sign(data []byte, privateKey []byte) ([]byte, error) {
	hash := sha256.New().Sum(data)

	key, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash)
}

func Verify(data []byte, publicKey []byte, sign []byte) error {
	hash := sha256.New().Sum(data)

	key, err := x509.ParsePKCS1PublicKey(publicKey)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hash, sign)
}
