package main

import "fmt"

func pkcs7(block []byte, size int) []byte  {
  n := len(block)
  p := size - (n % size)
  out := make([]byte, n+p);
  copy(out, block);
  for i:=0;i<p;i++{
    out = append(out, byte(0x04))
  }
  return out
}

func main() {
  fmt.Println(string(pkcs7([]byte("YELLOW SUBMARINE"), 20)))
}
