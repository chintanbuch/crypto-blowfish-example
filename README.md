# crypto-blowfish-example

Go Version ***1.7.5***

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