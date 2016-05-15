package util

import (
  "crypto/aes";
  "encoding/base64";
  "io/ioutil";
)

func AESECBDecrypt(ciphertext, key []byte) []byte  {
  block, err := aes.NewCipher(key)
  if err != nil {
    panic ("Unable to initialize aes cipher")
  }

  blockSize := block.BlockSize();

  if len(ciphertext)%blockSize != 0 {
    panic("Ciphertext size not proper");
  }

  plaintext := make([]byte, len(ciphertext))
  plaintextOriginal := plaintext

  for len(plaintext) > 0 {
    block.Decrypt(plaintext, ciphertext)
    plaintext = plaintext[blockSize:]
    ciphertext = ciphertext[blockSize:]
  }

  return plaintextOriginal
}

func Base64DecodeFile(filename string) []byte {
  contents, err := ioutil.ReadFile(filename);
  if err != nil {
    panic("Error reading file");
  }
  output, err1 := base64.StdEncoding.DecodeString(string(contents));
  if err1 != nil {
    panic("Decoding error");
  }
  return output
}
