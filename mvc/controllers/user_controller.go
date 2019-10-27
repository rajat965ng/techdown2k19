package controllers

import (
	"../services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(resp http.ResponseWriter, req *http.Request)  {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	if err!=nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}


	log.Println("The user id is: ", userId)

	user, err := services.GetUser(userId)
	if err!=nil {
		fmt.Println("Error: ",err)
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	buff,err := json.Marshal(user)
	resp.Write(buff)
}