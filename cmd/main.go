package main

import (
	"encoding/gob"
	"flag"
	activation_code "github.com/activation-code"
	"os"
	"time"
)

func main() {
	g := false
	flag.BoolVar(&g, "g", false, "生成")

	flag.Parse()

	flag.Usage()

	_ = activation_code.GenerateKeyPairFile("private.key", "public.key")

	lic := &activation_code.Licence{
		Product:   "iot-master",
		ExpireAt:  time.Now(),
		User:      "jason@zgwit.com",
		Signature: "0231687983216798516546898",
	}
	file, _ := os.OpenFile("gob", os.O_CREATE, os.ModePerm)
	enc := gob.NewEncoder(file)
	_ = enc.Encode(lic)
	file.Close()
}
