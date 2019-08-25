package api

import (
	"net/http"
	"stub/model"
	"stub/service"

	"github.com/gin-gonic/gin"
)

type LoginAPI struct {
	LoginService service.ILoginService
}

func (api LoginAPI) HandleFunc(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := api.LoginService.CheckLogin(username, password)
	statusCode := http.StatusOK
	StatusMessage := "ok"
	if err != nil {
		statusCode = http.StatusUnauthorized
		StatusMessage = err.Error()
	}
	c.JSON(statusCode, model.LoginResponse{
		Status: StatusMessage,
	})
}
