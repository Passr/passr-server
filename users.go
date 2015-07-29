package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"passr-server/authentication"
	"time"
	"strconv"
)

//-----------------------------------------------------------------
// User
//-----------------------------------------------------------------
type User struct {
	ID		   int    `json:"-" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Email              string
	Password	   string `json:"-" sql:"-"`
	EncryptedPassword  string `json:"-"`
	CreatedAt 	   time.Time
	UpdatedAt 	   time.Time
}

func (u *User) hashPassword() {
	u.EncryptedPassword = authentication.SaltAndHashPassword(u.Password)
}

func (u User) GetID() string {
	return strconv.Itoa(u.ID)
}

func (u *User) SetID(id string) (err error) {
	u.ID, err = strconv.Atoi(id)
	return
}

//-----------------------------------------------------------------
// POST /api/user
//-----------------------------------------------------------------
func createUser(c *gin.Context) {
	var newUser User

	if err := c.BindWith(&newUser, JsonApiBinding{}); err != nil {
		fmt.Println(err)

		// TODO: Reender JSON API error to client
		return
	}

	newUser.hashPassword()

	db.Create(&newUser)

	json, err := jsonapi.MarshalToJSON(newUser)	
	if err != nil {
		c.String(500, "Internal Server Error:"+err.Error())
		return
	}

	c.Data(201, CONTENT_TYPE, json)
}

//-----------------------------------------------------------------
// PUT /api/user
//-----------------------------------------------------------------

//-----------------------------------------------------------------
// DELETE /api/user
//-----------------------------------------------------------------
