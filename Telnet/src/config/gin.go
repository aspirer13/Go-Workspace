package config

import (
	"github.com/gin-gonic/gin"
)

// GinConfig stores configuration details for gin
type GinConfig struct {
	Mode string
	Host string
	Port int
}

//Gin represents gin set up
type Gin struct {
	Engine *gin.Engine
	Config GinConfig
}

func ginEngin() *Gin {
	var ginConfig GinConfig
	ginConfig.Host = "localhost"
	ginConfig.Port = 8088
	ginConfig.Mode = "debug"
	router := Gin{}
	router.Engine = gin.New()
	router.Config = ginConfig
	return &router
}
