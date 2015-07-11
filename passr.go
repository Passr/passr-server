package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwrd := []byte("SomeReallyLong Awful String")

	// Hashing the password givn witht he cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwrd, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Compare the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, passwrd)
	fmt.Println(err) // nil means  it is a match

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Passr.io")
	})

	r.Run(":3000")
}
