package app

import (
	"github.com/gin-gonic/gin"
	log "../log/logrus"
)


var (
	router *gin.Engine
)

func init()  {
	router = gin.Default()
}

func StartApp(){
	log.Info("about to map the urls", "step:1", "status:pending")
	mapUrls()
	log.Info("urls map successfully", "step:2", "status:successful")

	if err := router.Run(":9000"); err != nil {
		panic(err)
	}
}
