package service

import (
	"../domain/oauth"
	"github.com/pkg/errors"
	"time"
)

type oauthService struct {
}

type oauthService_interface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, error)
	GetAccessToken(token string) (*oauth.AccessToken, error)
}

var (
	OauthService oauthService_interface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, error := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if error != nil {
		return nil, error
	}
	token := &oauth.AccessToken{
		AccessToken: user.Username,
		UserId:      user.Id,
		Expires:     time.Now().UTC().Add(24 * time.Hour).Unix(),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}

	return token, error
}

func (s *oauthService) GetAccessToken(accesstoken string) (*oauth.AccessToken, error) {

	token, err := oauth.GetAccessTokenByToken(accesstoken)
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, errors.New("Token is expired")
	}
	return token, nil
}
