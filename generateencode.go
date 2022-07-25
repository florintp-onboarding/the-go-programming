// Fast multi-line base64 decode

package main

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/vault/helper/xor"
)

var encoded = "Value from $ENCODED"
var otp = "VAlue from $OTP"

func main() {
	tokenBytes, err := base64.RawStdEncoding.DecodeString(encoded)
	tokenBytes, err = xor.XORBytes(tokenBytes, []byte(otp))
	if err != nil {
		fmt.Println("Error")
		return
	}
	var err2 = string(tokenBytes)
	fmt.Println(err2)
}
