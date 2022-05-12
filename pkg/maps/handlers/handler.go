package maps

import (
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

func GetWorkspaces(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/workspaces/", ctx)
}

func CreateWorkspace(ctx *gin.Context) {
	util.PostGateWayUrl("http://localhost:3001/0.0/workspaces/workspace", ctx)
}

func GetWorkspace(ctx *gin.Context) {
	util.GetGateWayUrl("http://localhost:3001/0.0/workspaces/workspace", ctx)
}
