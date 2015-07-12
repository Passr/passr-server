package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"strconv"
	"time"
)

//----------------------------------------------------------------------------
// Credentials
//----------------------------------------------------------------------------
type Credential struct {
	ID               int    `json:"-" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	EncryptedBlobB64 string `sql:"type:text"`
	UserID           int    `json:"-"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (c Credential) GetID() string {
	return strconv.Itoa(c.ID)
}

func (c *Credential) SetID(id string) (err error) {
	c.ID, err = strconv.Atoi(id)
	return
}

//----------------------------------------------------------------------------
// GET /api/credentials
//----------------------------------------------------------------------------
func credentialsIndex(c *gin.Context) {
	var credentials []Credential

	db.Find(&credentials)

	json, err := jsonapi.MarshalToJSON(credentials)
	if err != nil {
		c.String(500, "Internal Server Error:"+err.Error())
		return
	}

	c.Data(200, "application/vnd.api+json", json)
}

//----------------------------------------------------------------------------
// GET /api/credentials/:id
//----------------------------------------------------------------------------
func credentialsShow(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	credential := &Credential{ID: id}

	query := db.First(credential)
	if query.Error != nil {
		if query.Error.Error() == "record not found" {
			c.String(404, "record not found")
			return
		}
	}

	json, err := jsonapi.MarshalToJSON(credential)
	if err != nil {
		c.String(500, "Internal Server Error:"+err.Error())
		return
	}

	c.Data(200, "application/vnd.api+json", json)
}

//----------------------------------------------------------------------------
// POST /api/credentials
//----------------------------------------------------------------------------
func credentialsCreate(c *gin.Context) {
	var newCredential Credential

	if err := c.BindWith(&newCredential, JsonApiBinding{}); err != nil {
		fmt.Println(err)

		// TODO: Render JSON API error to client
		return
	}

	db.Create(&newCredential)

	json, err := jsonapi.MarshalToJSON(newCredential)
	if err != nil {
		c.String(500, "Internal Server Error:"+err.Error())
		return
	}

	c.Data(201, CONTENT_TYPE, json)
}

//----------------------------------------------------------------------------
// PUT /api/credentials/:id
//----------------------------------------------------------------------------
func credentialsUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	savedCredential := &Credential{ID: id}

	query := db.First(savedCredential)
	if query.Error != nil {
		if query.Error.Error() == "record not found" {
			c.String(404, "record not found")
			return
		}
	}

	var intermCredential Credential
	if err := c.BindWith(&intermCredential, JsonApiBinding{}); err != nil {
		fmt.Println(err)
		return
	}

	savedCredential.EncryptedBlobB64 = intermCredential.EncryptedBlobB64

	query = db.Save(savedCredential)
	if query.Error != nil {
		fmt.Println(err)
		return
	}

	json, err := jsonapi.MarshalToJSON(savedCredential)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Data(200, "application/vnd.api+json", json)
}

//----------------------------------------------------------------------------
// DELETE /api/credentials/:id
//----------------------------------------------------------------------------
func credentialsDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	credential := &Credential{ID: id}

	query := db.Delete(credential)
	if query.Error != nil {
		if query.Error.Error() == "record not found" {
			c.String(404, "record not found")
			return
		}
	}

	c.String(204, "")
}
