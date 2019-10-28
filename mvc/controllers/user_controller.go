package controllers

import (
	"../services"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"../utils"
)

func GetUsers(c *gin.Context)  {

	userId, err := strconv.ParseInt(c.Param("user_id"),10,64)
	if err!=nil {
		utils.Respond(c,http.StatusBadRequest,errors.New("user_id must be a number").Error())
		return
	}


	log.Println("The user id is: ", userId)

	user, err := services.UserService.GetUser(userId)
	if err!=nil {
		fmt.Println("Error: ",err)
		utils.Respond(c,http.StatusNotFound,err.Error())
		return
	}
	utils.Respond(c,http.StatusOK,user)
}