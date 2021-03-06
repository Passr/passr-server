package main

import (
	"github.com/gin-gonic/gin"
	"passr-server/config"
)

// Temporarily set config path in constant
const (
	CONFIG_PATH = "config.conf"

	CONTENT_TYPE = "application/vnd.api+json"
)

var conf config.Config

func init() {
	var err error
	conf, err = config.Load(CONFIG_PATH)
	if err != nil {
		panic(err)
	}
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Passr.io")
	})

	// Credentials
	r.GET("/api/credentials", credentialsIndex)
	r.GET("/api/credentials/:id", credentialsShow)
	r.POST("/api/credentials", credentialsCreate)
	r.PUT("/api/credentials/:id", credentialsUpdate)
	r.DELETE("/api/credentials/:id", credentialsDelete)

	// User
	r.POST("/api/user", createUser)
	r.Run(":" + conf.Port)
}
