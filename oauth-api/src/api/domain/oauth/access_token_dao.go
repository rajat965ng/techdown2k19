package oauth

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	tokens = make(map[string]*AccessToken)
)

func (at *AccessToken) Save() (error) {
	accessToken := fmt.Sprintf("USR_%d", at.UserId)
	at.AccessToken = accessToken
	tokens[accessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, error) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.New("Invalid access token")
	}
	return token, nil
}
