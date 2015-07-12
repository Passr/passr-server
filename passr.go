package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"math/rand"
	"passr-server/config"
	"strconv"
	"time"
)

// Temporarily set config path in constant
const (
	CONFIG_PATH = "config.conf"

	CONTENT_TYPE = "application/vnd.api+json"
)

func main() {
	config, err := config.Load(CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Passr.io")
	})

	r.GET("/api/credentials", credentialsIndex)
	r.POST("/api/credentials", credentialsCreate)

	r.Run(":" + config.Port)
}

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

	c.Data(200, CONTENT_TYPE, json)
}
