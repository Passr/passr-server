package main

import (
	"github.com/gin-gonic/gin"
	"passr-server/config"
)

// Temporarily set config path in constant
const (
	CONFIG_PATH = "config.conf"
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

	r.Run(":" + config.Port)
}
