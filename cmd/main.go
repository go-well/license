package main

import (
	"flag"
	activation_code "github.com/activation-code"
)

func main() {
	g := false
	flag.BoolVar(&g, "g", false, "生成")

	flag.Parse()

	flag.Usage()

	_ = activation_code.GenerateKeyPairFile("private.key", "public.key")
}
