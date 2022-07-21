package main
import (
        "encoding/base64"
        "strings"
        "fmt"
        "os"
)
var  encoded = os.Args[1]
func main() {
split := strings.Split(encoded, ".")
for i := 0; i < len(split); i++ {
        tokenBytes, err := base64.RawStdEncoding.DecodeString(split[i])
        if err != nil {
          return
        }
        var sToken=string(tokenBytes)
        fmt.Printf("%s",sToken)
    }
}
