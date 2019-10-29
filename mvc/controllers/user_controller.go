package controllers

import (
	"../domains"
	"../services"
	"../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Access data source concurrently
func stackUsers(c *gin.Context, userBuff chan *domains.User, reqs []string) {
	for _, uid := range reqs {
		userId, err := strconv.ParseInt(uid, 10, 64)
		log.Println("UserId: ", userId)
		user, err := services.UserService.GetUser(userId)
		if err != nil {
			fmt.Println("Error: ", err)
			utils.Respond(c, http.StatusNotFound, err.Error())
			return
		}
		userBuff <- user
	}
	close(userBuff)
}

func GetUsers(c *gin.Context) {

	pathVariable := c.Param("user_id")
	log.Println("Path Variable: ", pathVariable)
	reqs := strings.Split(pathVariable, ",")
	log.Println("reqs: ", reqs)
	userBuff := make(chan *domains.User)
	users := make([]*domains.User, len(reqs))

	go stackUsers(c, userBuff, reqs) //Only sender can close the channel.

	x := 0
	for i := range userBuff {
		log.Println("User: ", i)
		users[x] = i
		x++
	}
	log.Println("users: ", users)
	utils.Respond(c, http.StatusOK, users)

}
