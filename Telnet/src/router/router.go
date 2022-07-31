package router

import (
	"fmt"
	"src/config"

	"github.com/gin-gonic/gin"
)

// dummy push

type Server struct {
	Gin *config.Gin
}

func (server *Server) Start() {
	gin.SetMode(gin.DebugMode)
	server.setUpRoutes()
	server.run()
}
func (server *Server) setUpRoutes() {
	routerGroup := server.Gin.Engine.Group("employee")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Employee App",
		})
	})
}
func (server *Server) run() {
	err := server.Gin.Engine.Run(fmt.Sprintf("%s:%v", server.Gin.Config.Host,
		server.Gin.Config.Port))
	if err != nil {
		panic(fmt.Sprintf("CRASH! Failed to start web server : %v", err.Error()))
	}
}
