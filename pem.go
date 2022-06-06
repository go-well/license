package activation_code

import (
	"encoding/pem"
	"os"
)

func WritePemFile(name, typ string, der []byte) error {
	block := &pem.Block{
		Type:  typ, //"RSA Private Key",
		Bytes: der,
	}

	file, err := os.OpenFile(name, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, block)
}

func ReadPemFile(name string, der []byte) ([]byte, error) {
	der, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(der)
	return block.Bytes, nil
}
