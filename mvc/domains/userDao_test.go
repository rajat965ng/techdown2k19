package domains

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDaoUserNotFound(t *testing.T)  {
	user,err := UserDao.GetUser(0)
	assert.Nil(t,user,"User with id 0 not found")
	assert.NotNil(t, err,"User with id 0 not found")
	assert.Error(t,err,"User not found with ID: 0")
}

func TestUserDaoUserFound(t *testing.T) {
	user,err:= UserDao.GetUser(1234)

	assert.NotNil(t,user,"User found with user id 1234")
	assert.Equal(t,user.FistName,"George","User has firstname:",user.FistName)
	assert.Nil(t,err,"There is no error")
}