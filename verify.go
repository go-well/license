package activation_code

import (
	"crypto/ed25519"
)

func Sign(data []byte, privateKey []byte) []byte {
	return ed25519.Sign(privateKey, data)
}

func Verify(data []byte, publicKey []byte, sign []byte) bool {
	return ed25519.Verify(publicKey, data, sign)
}
