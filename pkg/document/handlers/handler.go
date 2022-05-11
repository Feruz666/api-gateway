package document

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Example(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"resp": "HIIII",
	})
}
