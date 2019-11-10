package app

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
)

func Start_application() {
	router = gin.Default()
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
