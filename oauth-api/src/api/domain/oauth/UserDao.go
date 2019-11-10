package oauth

import "github.com/pkg/errors"

const (
	queryGetUserByUsernameAndPassword = "select id, username from users where username=? and password=?;"
)

var (
	users = map[string]*User{
		"fede": &User{Id: 123, Username: "fede",},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, error) {

	user := users[username]
	if user == nil {
		return nil, errors.New("No user with this username found")
	}
	return user, nil
}
