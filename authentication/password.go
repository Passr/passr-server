package authentication

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"encoding/base64"
)

const (
	SaltLength  = 64
	EncryptCost = 15
)

// This is returned when a new hash + salt combo is generated
type Password struct {
	// TODO: make hash []byte
	hash string
	salt string
}

func (p Password) hashB64() string {
	return base64.StdEncoding.EncodeToString([]byte(p.hash))
}

// This handles taking a raw user password and making it into something safe
// for storing in our DB
func hashPassword(salted_pass string) string {
	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(salted_pass), EncryptCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed_pass)
}

// Handles merging together the salt and the password
func combine(salt string, raw_pass string) string {

	// concat salt + password
	pieces := []string{salt, raw_pass}
	salted_password := strings.Join(pieces, "")
	return salted_password
}

// Generates a random salt using DevNull
func generateSalt() string {

	// Read in data
	data := make([]byte, SaltLength)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to a string
	salt := string(data[:])
	return salt
}

// Handles creating a new hash/salt combo from a raw password as
// the imput from the user
func CreatePassword(raw_pass string) *Password {

	password := new(Password)
	password.salt = generateSalt()
	salted_pass := combine(password.salt, raw_pass)
	password.hash = hashPassword(salted_pass)

	return password
}

func SaltAndHashPassword(raw_pass string) string {
	password := CreatePassword(raw_pass)
	pieces := []string{password.salt, password.hash}

	saltAndHash := strings.Join(pieces, ":")
        return base64.StdEncoding.EncodeToString([]byte(saltAndHash))
}

// Checks whether or not the correct password has been provided
func PasswordMatch(guess string, password *Password) bool {

	salted_guess := combine(password.salt, guess)

	// compare to the real password
	if bcrypt.CompareHashAndPassword([]byte(password.hash), []byte(salted_guess)) != nil {
		return false
	}

	return true
}
