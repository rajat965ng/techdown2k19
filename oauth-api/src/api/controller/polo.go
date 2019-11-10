package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Marco(c *gin.Context)  {
	c.String(http.StatusOK,"polo")
}