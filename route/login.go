package route

import (
	"stub/api"

	"github.com/gin-gonic/gin"
)

func NewServer(loginAPI *api.LoginAPI) *gin.Engine {
	router := gin.Default()
	router.POST("/login", loginAPI.HandleFunc)

	return router
}
