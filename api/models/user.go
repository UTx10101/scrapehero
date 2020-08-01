package models

import (
	// builtin
	"errors"
	"html"
	"log"
	"os"
	"strings"
	"time"

	// vendored
	"github.com/badoux/checkmail"
)

type User struct {
	Username   string
	Email      string
	Password   string
}

func (u *User) Prepare() {
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate() map[string]string {
	var errMessages = make(map[string]string)
	var err error

	if u.Password == "" {
		err = errors.New("Required Password")
		errorMessages["Required_password"] = err.Error()
	}
	if u.Email == "" {
		err = errors.New("Required Email")
		errorMessages["Required_email"] = err.Error()
	}
	
	return errMessages
}