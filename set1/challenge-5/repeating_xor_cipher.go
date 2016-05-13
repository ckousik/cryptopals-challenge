package main

import (
  "fmt";
  "encoding/hex";
)

func RepeatingXorCipher(plaintext []byte, key []byte) []byte  {
  n := len(plaintext)
  out := make([]byte, n)
  keysize := len(key)
  for i:= 0;i < n;i++{
    out[i] = plaintext[i] ^ key[i % keysize]
  }
  return out
}

func main() {
  input, key := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"), []byte("ICE")
  output := RepeatingXorCipher(input,key)
  fmt.Printf("%s\n", hex.EncodeToString(output))
}
