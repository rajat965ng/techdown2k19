package oauth

import (
	"github.com/pkg/errors"
	"strings"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() error {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.New("Invalid Username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.New("Invalid Password")
	}
	return nil
}
