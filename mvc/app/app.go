package app

import (
	uc "../controllers"
	"net/http"
)


func StartApp(){

	http.HandleFunc("/users",uc.GetUsers)

	if err := http.ListenAndServe(":9000",nil); err != nil {
		panic(err)
	}
}
