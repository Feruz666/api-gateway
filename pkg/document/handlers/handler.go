package document

import (
	"net/http"

	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

func Example(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"resp": "HIIII",
	})
}

func GetExample(ctx *gin.Context) {
	util.GetGateWayUrl("https://jsonplaceholder.typicode.com/posts", ctx)
}



func DeleteExample(ctx *gin.Context) {
	util.DeleteGateWayUrl("https://jsonplaceholder.typicode.com/posts", ctx)
}