package main

import (
    "./chintanbuch.com/example/gocrypto/util"
    "fmt"
    "os"
)

/*
    Go Version: 1.7.5s
    1) go get golang.org/x/crypto/blowfish
    2) ./build.sh
    3) command to encrypt: "./run.sh test"
        -- result --
        Encrypting "test"
        BlowFish Encrypted String: AAAAAAAAAABYm+3y9AZZzQ==
    4) command to decrypt: "./run.sh AAAAAAAAAABYm+3y9AZZzQ== 1"
        -- result --
        Decrypting "AAAAAAAAAABYm+3y9AZZzQ=="
        BlowFish Decrypted String: test
*/

func main() {
    // default string to encrypt
    var strToEncrDcr string
    var decryptOnly bool = false
    // check for arguments if any
    strArg := os.Args[1:]
    if (len(strArg) == 1) {
        strToEncrDcr = strArg[0]
        fmt.Println("Encrypting \"" + strToEncrDcr + "\"")
    } else if (len(strArg) == 2) {
        if (strArg[1] != "1") {
            fmt.Println("Wrong Inputs!")
        }
        strToEncrDcr = strArg[0]
        decryptOnly = true
        fmt.Println("Decrypting \"" + strToEncrDcr + "\"")
    }

    if (decryptOnly) {
        decr, err := util.DecryptToString(strToEncrDcr)
        if err != nil { fmt.Println("Error Decrypting!") } else {fmt.Println("BlowFish Decrypted String: " + decr)}
    } else {
        encr, err := util.EncryptToString(strToEncrDcr)
        if err != nil { fmt.Println("Error Encrypting!") } else {fmt.Println("BlowFish Encrypted String: " + encr)}
    }
}