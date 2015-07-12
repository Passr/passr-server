package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"math/rand"
	"strconv"
	"time"
)

//----------------------------------------------------------------------------
// Credentials
//----------------------------------------------------------------------------
type Credential struct {
	ID               string `json:"-"`
	EncryptedBlobB64 string
	UserID           int `json:"-"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (c Credential) GetID() string {
	return c.ID
}

func (c *Credential) SetID(id string) error {
	c.ID = id
	return nil
}

var credentials []Credential = []Credential{}

//----------------------------------------------------------------------------
// GET /api/credentials
//----------------------------------------------------------------------------
func credentialsIndex(c *gin.Context) {
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
	id := c.Param("id")

	for _, x := range credentials {
		if x.ID == id {
			json, err := jsonapi.MarshalToJSON(x)
			if err != nil {
				c.String(500, "Internal Server Error:"+err.Error())
				return
			}

			c.Data(200, "application/vnd.api+json", json)

			return
		}
	}

	c.String(404, "NOT FOUND")
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

	newCredential.ID = strconv.Itoa(rand.Int())

	credentials = append(credentials, newCredential)

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
	id := c.Param("id")

	for _, x := range credentials {
		if x.ID == id {
			if err := c.BindWith(&x, JsonApiBinding{}); err != nil {
				fmt.Println(err)

				// TODO: Render JSON API error to client
				return
			}

			x.ID = id

			json, err := jsonapi.MarshalToJSON(x)
			if err != nil {
				c.String(500, "Internal Server Error:"+err.Error())
				return
			}

			c.Data(200, "application/vnd.api+json", json)

			return
		}
	}

	c.String(404, "NOT FOUND")
}
