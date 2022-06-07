package license

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/super-l/machine-code/machine"
	"time"
)

type Licence struct {
	Product   string `json:"product"`
	User      string `json:"user"`
	UUID      string `json:"UUID,omitempty"`
	SN        string `json:"SN,omitempty"`
	CPUID     string `json:"CPUID,omitempty"`
	MAC       string `json:"MAC,omitempty"`
	ExpireAt  string `json:"expireAt"`
	Signature string `json:"signature"`
}

func (l *Licence) Match(product string) error {
	if l.Product != product {
		return fmt.Errorf("产品 不匹配，证书 %s 本机 %s", l.Product, product)
	}

	md := machine.GetMachineData()
	if l.UUID != "" && l.UUID != md.PlatformUUID {
		return fmt.Errorf("UUID 不匹配，证书 %s 本机 %s", l.UUID, md.PlatformUUID)
	}

	if l.SN != "" && l.SN != md.SerialNumber {
		return fmt.Errorf("SN 不匹配，证书 %s 本机 %s", l.SN, md.SerialNumber)
	}

	if l.CPUID != "" && l.CPUID != md.CpuId {
		return fmt.Errorf("CPUID 不匹配，证书 %s 本机 %s", l.CPUID, md.CpuId)
	}

	if l.MAC != "" && l.MAC != md.Mac {
		return fmt.Errorf("MAC 不匹配，证书 %s 本机 %s", l.MAC, md.Mac)
	}

	tm, err := time.Parse("2006-01-02", l.ExpireAt)
	if err != nil {
		return err
	}

	if tm.Before(time.Now()) {
		return fmt.Errorf("证书已经失效 %s", tm.Format("2006-01-02"))
	}

	return nil
}

func (l *Licence) Sign(privateKey []byte) {
	str := l.Product + l.User + l.UUID + l.SN + l.CPUID + l.MAC + l.ExpireAt
	sign := ed25519.Sign(privateKey, []byte(str))
	l.Signature = base64.StdEncoding.EncodeToString(sign)
}

func (l *Licence) Verify(publicKey []byte) error {
	str := l.Product + l.User + l.UUID + l.SN + l.CPUID + l.MAC + l.ExpireAt
	sign, err := base64.StdEncoding.DecodeString(l.Signature)
	if err != nil {
		return err
	}
	if !ed25519.Verify(publicKey, []byte(str), sign) {
		return fmt.Errorf("无效证书")
	}
	return nil
}

func (l *Licence) Encode() string {
	text, _ := json.Marshal(l)
	return base64.StdEncoding.EncodeToString(text)
}

func (l *Licence) Decode(text string) error {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, l)
}
