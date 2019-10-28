package services

import (
	"../domains"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	UserDaoMock userDaoMock
	getUserFunction func(userId int64) (*domains.User,error)
)


type userDaoMock struct {}


func (u *userDaoMock) GetUser(userId int64) (*domains.User,error){
	return getUserFunction(userId)
}

func init()  {
	domains.UserDao = &UserDaoMock
}

func TestUserService_GetUserNotFoundInDB(t *testing.T) {

	getUserFunction = func(userId int64) (*domains.User, error) {
		return nil,errors.New("User 0 does not exist")
	}
	user,err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)

}

func TestUserService_GetUserFoundInDB(t *testing.T) {

	getUserFunction = func(userId int64) (*domains.User, error) {
		return &domains.User{Id:1234},nil
	}
	user,err := UserService.GetUser(1234)

	assert.NotNil(t, user)
	assert.Nil(t, err)

}
