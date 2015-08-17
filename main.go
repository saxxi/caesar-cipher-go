package main

import (
  "fmt"
  "github.com/saxxi/caesar-cipher/CaesarCipher"
)

func main() {

  encoded := CaesarCipher.Encode("汉语 Aa - Hi How are you uvxyz?", 3)

  fmt.Println(encoded)
  fmt.Println(CaesarCipher.Decode(encoded, 3))

}
