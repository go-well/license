package activation_code

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

func GenerateKeyPair(bits int) ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := privateKey.PublicKey

	privateDer := x509.MarshalPKCS1PrivateKey(privateKey)
	publicDer := x509.MarshalPKCS1PublicKey(&publicKey)

	return privateDer, publicDer, nil
}

func GenerateKeyPairFile(bits int, private, public string) error {
	privateDer, publicDer, err := GenerateKeyPair(bits)
	if err != nil {
		return err
	}

	err = WritePemFile(private, "RSA Private Key", privateDer)
	if err != nil {
		return err
	}
	return WritePemFile(public, "RSA Public Key", publicDer)
}
