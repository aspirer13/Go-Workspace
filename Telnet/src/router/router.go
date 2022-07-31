package router

import (
	"fmt"
	"net/http"
	"src/config"
	"src/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *config.Gin
}

func (server *Server) Start() {
	gin.SetMode(gin.DebugMode)
	server.setUpRoutes()
	server.run()
}
func (server *Server) setUpRoutes() {
	routerGroup := server.Gin.Engine.Group("telnet")
	routerGroup.GET("/deleteAllData", func(c *gin.Context) {
		service.DeleteAllRecords()
		c.JSON(http.StatusOK, gin.H{"data": "All Data Deleted"})
	})
	routerGroup.GET("/ingestplans", func(c *gin.Context) {
		service.IngestPlansFromCsvFile()
		c.JSON(http.StatusOK, gin.H{"data": "All plans added"})
	})
	routerGroup.GET("/ingestConsumers", func(c *gin.Context) {
		service.InjestConsumersFromCsvFile()
		c.JSON(http.StatusOK, gin.H{"data": "All consumers added"})
	})
}
func (server *Server) run() {
	err := server.Gin.Engine.Run(fmt.Sprintf("%s:%v", server.Gin.Config.Host,
		server.Gin.Config.Port))
	if err != nil {
		panic(fmt.Sprintf("CRASH! Failed to start web server : %v", err.Error()))
	}
}
