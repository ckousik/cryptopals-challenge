package main

import (
  "fmt";
  "encoding/hex";
)

func byteXor(k byte, src []byte) []byte  {
  plain := make([]byte, len(src))
  for i := range src {
    plain[i] = src[i] ^ k
  }
  return plain
}

func freqScore(plaintext []byte) uint {
  var score uint
  for _, i := range plaintext {
    score += letterScore(i)
  }
  return score
}

func letterScore(l byte) uint {
  freq := []byte("ETAOINSHRDLU etaoinshrdlu")
  for _, c := range freq {
    if(c == l){
      return 3
    }
  }
  if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') {
    return 2
  }
  if l >= '0' && l <='9' {
    return 1
  }
  return 0
}

func SingleByteXor(cipher []byte) (key byte) {
  var score uint
  for k := 0x00; k <= 0xff; k++ {
    plain := byteXor(byte(k), cipher)
    if s:= freqScore(plain); s > score {
      key = byte(k)
      score = s
    }
  }
  return
}

func main() {
  input, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
  k := SingleByteXor(input)
  fmt.Printf("Key:%s Output:%s\n", string(k), string(byteXor(k ,input)))
}
