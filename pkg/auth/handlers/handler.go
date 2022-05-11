package handlers

import (
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	util.GateWayUrl("http://localhost:8080/signup", ctx)
}

func LoginUser(ctx *gin.Context) {
	util.GateWayUrl("http://localhost:8080/users/login", ctx)
}
