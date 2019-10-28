package app

import uc "../controllers"

func mapUrls()  {
	router.GET("/users/:user_id",uc.GetUsers)
}
