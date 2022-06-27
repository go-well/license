package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	activation_code "github.com/zgwit/go-license"
	"github.com/super-l/machine-code/machine"
)

func main() {
	g := false
	flag.BoolVar(&g, "g", false, "生成")

	flag.Parse()

	flag.Usage()

	pub, key, _ := activation_code.GenerateKeyPair()
	fmt.Println(hex.EncodeToString(pub))
	fmt.Println(hex.EncodeToString(key))

	md := machine.GetMachineData()

	lic := &activation_code.Licence{
		Product:  "iot-master",
		User:     "jason",
		UUID:     md.PlatformUUID,
		SN:       md.SerialNumber,
		CPUID:    md.CpuId,
		MAC:      md.Mac,
		ExpireAt: "2022-06-08",
		//Signature: "0231687983216798516546898",
	}

	lic.Sign(key)
	fmt.Println(lic.Encode())

	err := lic.Verify(pub)
	if err != nil {
		panic(err)
	}
	err = lic.Match("iot-master")
	if err != nil {
		panic(err)
	}
}
