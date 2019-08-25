package service_test

import (
	. "stub/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckLoginFailedWithNotExistingUsername(t *testing.T) {
	loginService := &LoginServiceImpl{}
	loginStatus, err := loginService.CheckLogin("fakeusername", "noempty")

	assert.False(t, loginStatus)
	assert.EqualError(t, err, "not existing username")
}

func TestCheckLoginFailedWithWrongPassword(t *testing.T) {
	loginService := &LoginServiceImpl{}
	loginStatus, err := loginService.CheckLogin("imkk-000", "wrongP@ssw0rd")

	assert.False(t, loginStatus)
	assert.EqualError(t, err, "wrong password")
}

func TestCheckLoginPassed(t *testing.T) {
	loginService := &LoginServiceImpl{}
	loginStatus, _ := loginService.CheckLogin("imkk-000", "1mkkn@ja*")

	assert.True(t, loginStatus)
}
