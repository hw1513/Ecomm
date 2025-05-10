package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hw1513/ecomm/common/config"
)

func RunRESTServer(config *config.ServiceConfig, wrapper func(router *gin.Engine)) {
	addr := config.HTTPAddr
	if addr == "" {
		//warn
	}

	apiRouter := gin.New()
	wrapper(apiRouter)
	apiRouter.Group("/api")
	if err := apiRouter.Run(addr); err != nil {
		panic(err)
	}
}
