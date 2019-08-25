package service_test

import (
	. "stub/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLoginFailedWithNotExistingUsername(t *testing.T) {
	loginService := &LoginService{}
	err := loginService.CheckLogin("fakeusername", "noempty")

	assert.EqualError(t, err, "not existing username")
}

func TestCheckLoginFailedWithWrongPassword(t *testing.T) {
	loginService := &LoginService{}
	err := loginService.CheckLogin("imkk-000", "wrongP@ssw0rd")

	assert.EqualError(t, err, "wrong password")
}

func TestCheckLoginPassed(t *testing.T) {
	loginService := &LoginService{}
	err := loginService.CheckLogin("imkk-000", "1mkkn@ja*")

	assert.Nil(t, err)
}
