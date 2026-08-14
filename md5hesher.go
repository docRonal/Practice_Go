package main
import (
"crypto/md5"
"encoding/hex"
"fmt"
)

func main() {
targetHash := "81dc9bdb52d04dc20036dbd8313ed055"


for i:=0; i<100000;i++{
pin := fmt.Sprintf("%04d", i)

hasher := md5.New()
hasher.Write([]byte(pin))
hashString := hex.EncodeToString(hasher.Sum(nil))

//fmt.Println(hashString, "=", pin)

if targetHash==hashString {
	fmt.Println("your pin is: ", pin)
	break
}


}

}
