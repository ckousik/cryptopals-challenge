package main

import (
  "fmt";
  "io/ioutil";
  "encoding/base64";
)

func byteXor(k byte, src []byte) []byte  {
  plain := make([]byte, len(src))
  for i := range src {
    plain[i] = src[i] ^ k
  }
  return plain
}

func freqScore(plaintext []byte) int {
  var score int
  for _, i := range plaintext {
    score += letterScore(i)
  }
  return score
}

func letterScore(l byte) int {
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
  var score int
  for k := 0x00; k <= 0xff; k++ {
    plain := byteXor(byte(k), cipher)
    if s:= freqScore(plain); s > score {
      key = byte(k)
      score = s
    }
  }
  return
}

func HammingDistance(a, b []byte) int{
  if(len(a) != len(b)){
    panic("Different length strings")
  }
  var dist int
  dist = 0
  for i:=0; i < len(a); i++ {
    c := a[i] ^ b[i]
    for c != 0 {
      dist += int(c & byte(0x01))
      c = c >> 1
    }
  }
  return dist
}

func BreakRepeatedXor(cipher []byte) []byte {

  keySize := GetBestKeySize(cipher)
  key := []byte{}
  for _,part := range Transpose(cipher, keySize){
    key = append(key,SingleByteXor(part))
  }

  return key
}

func GetBestKeySize(cipher []byte) int{
  best, result, checks  := 0, 0, len(cipher)/40
  for keySize:=2;keySize<40;keySize++ {
    dist := 0
    first := cipher[:keySize]
    for i:=1; i< checks;i++ {
      next := cipher[keySize*i : keySize*(i+1)]
      dist += HammingDistance(first, next)
    }
    dist /= keySize
    if best ==0 || dist < best {
      best = dist
      result = keySize
    }
  }
  return result
}

func Transpose(cipher []byte, size int) [][]byte{
  result := make([][]byte, size)
  i := 0
  for _, j := range cipher {
    result[i] = append(result[i],j)
    i++
    if i == size {
      i=0
    }
  }
  return result
}

func DecodeXor(cipher, key []byte) []byte  {
  plain := make([]byte, len(cipher))
  ks := len(key)
  for i := range cipher {
    plain[i] = cipher[i] ^ key[i % ks]
  }
  return plain
}

func main() {
  input, _ := ioutil.ReadFile("6.txt")
  input, _ = base64.StdEncoding.DecodeString(string(input))
  key := BreakRepeatedXor(input)
  fmt.Printf("Output:%sKey:%s\n", string(DecodeXor(input,key)),string(key))
}
