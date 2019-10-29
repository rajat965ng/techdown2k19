package domains

import (
	"errors"
	"fmt"
	"log"
)

var (
	users = map[int64]*User{

		123: { Id:123, FistName:"Tobby", LastName:"Mcguire", Email:"tobby@gmail.com" },
		1234: { Id:1234, FistName:"George", LastName:"Pool", Email:"myemail@gmail.com" },
	}

	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(userId int64) (*User,error)
}


func init()  {
	UserDao = &userDao{}
}


type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User,error) {

	log.Println("We are accessing the DB.")

	if user := users[userId]; user!=nil  {
		return user,nil
	}
	return  nil,errors.New(fmt.Sprintf("User not found with ID: %v",userId))
}