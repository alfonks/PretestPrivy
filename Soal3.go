package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "math/big"
    "fmt"
    "io"
    "log"
)

func encrypt(text []byte, key []byte) []byte{ //encryp text function

	c, err := aes.NewCipher(key) //membuat cipher dengan key yang sudah dibuat
    
    if err != nil { //jika terdapat error maka akan mengeluarkan log error tersebut
        log.Fatal(err)
    }

	gcm, err := cipher.NewGCM(c) //membuat GCM berdasarkan cipher diatas
    
    if err != nil { //jika terdapat error maka akan mengeluarkan log error tersebut
        log.Fatal(err)
    }

    nonce := make([]byte, gcm.NonceSize()) //membuat array byte berdasarkan panjang gcm

    if _, err = io.ReadFull(rand.Reader, nonce); err != nil { //jika terdapat error maka akan mengeluarkan log error tersebut
        log.Fatal(err)
    }

   return gcm.Seal(nonce, nonce, text, nil) //men encrypt text 
}

func decrypt(text []byte, key []byte) string{ //fungsi untuk men decrypt text

    c, err := aes.NewCipher(key) //membuat cipher berdasarkan key

    if err != nil { //jika terdapat error maka akan mengeluarkan log error tersebut
        log.Fatal(err)
    }

    gcm, err := cipher.NewGCM(c) //Membuat GCM berdasarkan cipher diatas

    if err != nil { //jika terdapat error maka akan mengeluarkan log error tersebut
        log.Fatal(err)
    }

    nonceSize := gcm.NonceSize() //menyimpan panjang gcm

    if len(text) < nonceSize { //jika ukuran text lebih kecil dari nonceSize maka akan mengeluarkan log error
        log.Fatal(err)
    }

    nonce, text := text[:nonceSize], text[nonceSize:] 
    plaintext, err := gcm.Open(nil, nonce, text, nil) //men-decrypt text yang tadinya di encrypt

    if err != nil { //jika terdapat error maka akan mengeluarkan log error tersebut
        log.Fatal(err)
    }

    return string(plaintext) //mereturn text yang sudah di decrpyt
}

func increaseKeySize(key []byte) []byte {

	for i := len(key) ; i < 32 ; i++{ //loop selama size key lebih kecil dari 32 maka akan menambahkan random integer
		newInteger, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
            log.Fatal(err)
		}
        byteinteger := newInteger.Int64()
        key = append(key, byte(byteinteger))
		
	}
    fmt.Println("")
	return key
}

func main(){
	text := []byte("Selamat Datang")
	key := []byte("1234privy5678")

	key = increaseKeySize(key) //Menaikan Key size dari 13 byte jadi 32 byte karena AES 256 membutuhkan 32byte key length
    fmt.Printf("New key: %v \n", string(key))


	encryptedtext := encrypt(text, key)
	fmt.Println("Encrypted Text:")
	fmt.Printf("%v\n\n",encryptedtext)

	decryptedtext := decrypt(encryptedtext, key)
	fmt.Println("Decrypted Text:")
	fmt.Printf("%s\n\n",decryptedtext)
}