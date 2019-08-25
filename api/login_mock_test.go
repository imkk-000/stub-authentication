package api_test

import (
	"errors"
)

type MockLoginServicePassed struct {}
func (MockLoginServicePassed) CheckLogin(username, password string) error {
	return nil
}

type MockLoginServiceFailedWithWrongPassword struct {}
func (MockLoginServiceFailedWithWrongPassword) CheckLogin(username, password string) error {
	return errors.New("wrong password")
}

type MockLoginServiceFailedWithNotExistingUsername struct {}
func (MockLoginServiceFailedWithNotExistingUsername) CheckLogin(username, password string) error {
	return errors.New("not existing username")
}
