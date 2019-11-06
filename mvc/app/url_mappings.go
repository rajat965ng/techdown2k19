package app

import uc "../controllers"
import log "../log/zap"
import "go.uber.org/zap"

func mapUrls()  {
	log.Info("about to map the urls", zap.Any("step",3), zap.Any("status","pending"))

	router.GET("/users/:user_id",uc.GetUsers)
	log.Info("urls map successfully", zap.Any("step",4), zap.Any("status","completed"))

}
