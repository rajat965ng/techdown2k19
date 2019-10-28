package services

import "../domains"

type userService struct {}

var (
	UserService userService
)

func (u *userService) GetUser(userid int64) (user *domains.User, err error ){
	return domains.UserDao.GetUser(userid)
}