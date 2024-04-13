package deb

import (
	"os/user"
)

func CurrentUserDir() string {
	currUser, err := user.Current()
	if err != nil {
		log.Errorf("cannot get current user default path, %s", err)
		return ""
	}
	userPath := currUser.HomeDir
	return userPath
}
