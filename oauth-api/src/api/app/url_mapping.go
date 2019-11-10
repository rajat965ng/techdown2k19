package app

import (
	"../controller"
	"../controller/oauth"
)

func mapUrls() {

	router.GET("/marco", controller.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
