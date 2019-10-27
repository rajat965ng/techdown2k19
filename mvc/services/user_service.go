package services

import "../domains"

func GetUser(userid int64) (user *domains.User, err error ){
	return domains.UserDao(userid)
}