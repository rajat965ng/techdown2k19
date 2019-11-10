package oauth

import (
	"github.com/gin-gonic/gin"
	"../../domain/oauth"
	"net/http"
	"../../service"
)

var (
	request oauth.AccessTokenRequest
)

func CreateAccessToken(c *gin.Context) {

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}
	token, err := service.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")
	token, err := service.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
