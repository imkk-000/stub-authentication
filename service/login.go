package service

import "errors"

type LoginService interface {
	CheckLogin(username, password string) error
}

type LoginServiceImpl struct{}

const fakeUsername = "imkk-000"
const fakePassword = "1mkkn@ja*"

func (LoginServiceImpl) CheckLogin(username, password string) error {
	if username != fakeUsername && username != fakePassword {
		return errors.New("not existing username")
	}
	if username == fakeUsername {
		if password != fakePassword {
			return errors.New("wrong password")
		} else {
			return nil
		}
	}
	return errors.New("unknow types")
}
