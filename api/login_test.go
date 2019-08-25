package api_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	. "stub/api"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginAPIHandlerWithLoginFailedNotExistingUsername(t *testing.T) {
	reqBody := "username=fakeusername&password=noempty"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	rec := httptest.NewRecorder()
	loginAPI := LoginAPI{
		LoginService: &MockLoginServiceFailedWithNotExistingUsername{},
	}
	router := gin.Default()
	router.POST("/login", loginAPI.HandleFunc)

	router.ServeHTTP(rec, req)
	respBody, _ := ioutil.ReadAll(rec.Body)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, []byte(`{"status":"not existing username"}`), respBody)
}

func TestLoginAPIHandlerWithLoginFailedWrongPassword(t *testing.T) {
	reqBody := "username=imkk-000&password=wrongP@ssw0rd"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	rec := httptest.NewRecorder()
	loginAPI := LoginAPI{
		LoginService: &MockLoginServiceFailedWithWrongPassword{},
	}
	router := gin.Default()
	router.POST("/login", loginAPI.HandleFunc)

	router.ServeHTTP(rec, req)
	respBody, _ := ioutil.ReadAll(rec.Body)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Equal(t, []byte(`{"status":"wrong password"}`), respBody)
}

func TestLoginAPIHandlerWithLoginPassed(t *testing.T) {
	reqBody := "username=imkk-000&password=1mkkn@ja*"
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	rec := httptest.NewRecorder()
	loginAPI := LoginAPI{
		LoginService: &MockLoginServicePassed{},
	}
	router := gin.Default()
	router.POST("/login", loginAPI.HandleFunc)

	router.ServeHTTP(rec, req)
	respBody, _ := ioutil.ReadAll(rec.Body)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, []byte(`{"status":"ok"}`), respBody)
}
