package main

import (
  "fmt";
  "../../util";
)

func main() {
  input := util.Base64DecodeFile("7.txt")
  output := util.AESECBDecrypt(input, []byte("YELLOW SUBMARINE"))
  fmt.Println(string(output))
}
