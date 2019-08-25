package service

import "errors"

type LoginService interface {
	CheckLogin(username, password string) (bool, error)
}

type LoginServiceImpl struct{}

const fakeUsername = "imkk-000"
const fakePassword = "1mkkn@ja*"

func (LoginServiceImpl) CheckLogin(username, password string) (bool, error) {
	if username != fakeUsername && username != fakePassword {
		return false, errors.New("not existing username")
	}
	if username == fakeUsername {
		if password != fakePassword {
			return false, errors.New("wrong password")
		} else {
			return true, nil
		}
	}
	return false, errors.New("unknow types")
}
