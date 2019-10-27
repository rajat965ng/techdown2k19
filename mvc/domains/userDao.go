package domains

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{

		1234: { Id:1234, FistName:"George", LastName:"Pool", Email:"myemail@gmail.com" },
	}
)

func UserDao(userId int64) (*User,error) {
	if user := users[userId]; user!=nil  {
		return user,nil
	}
	return  nil,errors.New(fmt.Sprintf("User not found with ID: %v",userId))
}