package auth

import (
	"log"

	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

var url = GetAuthConfig()

func CreateUser(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/users", ctx)
}

func LoginUser(ctx *gin.Context) {
	util.PostGateWayUrl(url+"/users/login", ctx)
}

func GetAuthConfig() string {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	return config.AUTH_SYSTEM_ADDRESS
}
