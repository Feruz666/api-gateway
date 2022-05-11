package auth

import (
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	util.PostGateWayUrl("http://localhost:8080/signup", ctx)
}

func LoginUser(ctx *gin.Context) {
	util.PostGateWayUrl("http://localhost:8080/users/login", ctx)
}
