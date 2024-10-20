package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/argon2"
)

type Argon2Hash struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
	saltLen uint32
}

type HashSalt struct {
	Hash          []byte
	salt          []byte
	StringToStore string
}

// Initializer for algorithm
func NewArgon2Hash(time, keyLen, saltLen, memory uint32, threads uint8) *Argon2Hash {
	return &Argon2Hash{
		time:    time,
		keyLen:  keyLen,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
	}
}

// Generating salt (byte slice with pseudorandom symbols)
func SaltGen(length uint32) ([]byte, error) {
	salt := make([]byte, length)

	// Read pseudorandom numbers and write them in salt
	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}

	return salt, nil

}

// Generating hash with given password and salt
func (a *Argon2Hash) HashGen(password, salt []byte) (*HashSalt, error) {
	var err error

	// If salt len is zero => we hashing new password => we need new salt ;)
	if len(salt) == 0 {
		salt, err = SaltGen(a.saltLen)
	}

	if err != nil {
		return nil, err
	}

	// Generating password hash with given params
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)

	// Salt and hash is byte slice and to properly read them (as string) we need to use this functions
	saltDecoded := base64.RawStdEncoding.EncodeToString(salt)
	hashDecoded := base64.RawStdEncoding.EncodeToString(hash)

	// Creating string with salt and password hash
	stringToStore := fmt.Sprintf("$m=%d,t=%d$%s$%s", a.memory, a.time, saltDecoded, hashDecoded)

	// This is unnecessary structure i created following the guide 0_0
	return &HashSalt{Hash: hash, salt: salt, StringToStore: stringToStore}, nil

}

// Returns error if password hashes are not equal
func (a *Argon2Hash) Compare(hash, salt, password []byte) error {

	// We hash given password with salt we created for original one
	// so we can check wether hashes are equal
	hashSalt, err := a.HashGen(password, salt)

	if err != nil {
		return err
	}

	// Comparing hashes
	if !bytes.Equal(hash, hashSalt.Hash) {
		return fmt.Errorf("invalid password (hashes does not match)")
	}

	return nil
}

// Function to extract hash and salt from password string
// that is stored in db
//
// returns hash, salt, error
func DecodeArgon2String(passwordHash string) ([]byte, []byte, error) {
	values := strings.Split(passwordHash, "$")

	// Transform string hash in byte slice
	salt, err := base64.RawStdEncoding.Strict().DecodeString(values[2])
	if err != nil {
		return nil, nil, err
	}

	// Transform string hash in byte slice
	hash, err := base64.RawStdEncoding.Strict().DecodeString(values[3])
	if err != nil {
		return nil, nil, err
	}

	return hash, salt, nil
}

var Argon2Hasher = NewArgon2Hash(2, 100, 32, 64*1024, 1)

func Test() {
	pas := []byte("qwerty1")

	testDbString := "$m=65536,t=2$u0h2MoT48NXJFIRQXW9/i0tNDu427RJv3vIeZQIm8FU$QtUmjmPhsBgWGloNQVRoFkLkHnQwuCqRVfgKA0Sm2QNMPc86vSLxQ/c8JUXroO37qwXdfC9DNvTnm/OOi7GfTBGnJJotLBlonG/9epAMGY453s9UZVeghmvftCagXzPS9QB3cg"

	var salt []byte

	cringe, err := Argon2Hasher.HashGen(pas, salt)

	if err != nil {
		return
	}

	storedH, storedS, err := DecodeArgon2String(testDbString)
	if err != nil {
		fmt.Println("CRINGE", err)
		return
	}

	err = Argon2Hasher.Compare(storedH, storedS, []byte("qwerty1"))
	if err != nil {
		fmt.Println("DONT MATCH")
	} else {
		fmt.Println("MATCH")
	}

	if err := godotenv.Load("secrets.env"); err != nil {
		fmt.Println("Error load .env")
	}

	// fmt.Println(testDbString)
	fmt.Printf("%s", cringe.StringToStore)
	fmt.Print("\n\n\n\n")
}
