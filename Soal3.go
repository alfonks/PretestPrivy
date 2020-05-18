package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "math/big"
    "fmt"
    "io"
)

func encrypt(text []byte, key []byte) []byte{

	c, err := aes.NewCipher(key)
    
    if err != nil {
        fmt.Println(err)
    }

	gcm, err := cipher.NewGCM(c)
    
    if err != nil {
        fmt.Println(err)
    }

    nonce := make([]byte, gcm.NonceSize())

    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        fmt.Println(err)
    }

   return gcm.Seal(nonce, nonce, text, nil)
}

func decrypt(text []byte, key []byte) string{

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        fmt.Println(err)
    }

    nonceSize := gcm.NonceSize()
    if len(text) < nonceSize {
        fmt.Println(err)
    }

    nonce, text := text[:nonceSize], text[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, text, nil)

    if err != nil {
        fmt.Println(err)
    }

    return string(plaintext)
}

func increaseKeySize(key []byte) []byte {

	for i := len(key) ; i < 32 ; i++{
		newInteger, err := rand.Int(rand.Reader, big.NewInt(100))
		if err == nil {
			byteinteger := newInteger.Int64()
			fmt.Println(byteinteger)
			key = append(key, byte(byteinteger))
		}
		
	}
	return key
}

func main(){
	text := []byte("Selamat Datang")
	key := []byte("1234privy5678")

	key = increaseKeySize(key) //Menaikan Key size dari 13 byte jadi 32 byte 

	encryptedtext := encrypt(text, key)
	fmt.Println("Encrypted Text:")
	fmt.Printf("%v\n\n",encryptedtext)

	decryptedtext := decrypt(encryptedtext, key)
	fmt.Println("Decrypted Text:")
	fmt.Printf("%s\n\n",decryptedtext)
}