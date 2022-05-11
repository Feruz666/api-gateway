package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GateWayUrl(url string, ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponce(err))
	}

	responseBody := bytes.NewBuffer(data)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, ErrorResponce(err))
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponce(err))
	}

	jsonData := make(map[string]interface{})
	json.Unmarshal([]byte(string(b)), &jsonData)

	ctx.JSON(resp.StatusCode, jsonData)
}

// ErrorResponce ...
func ErrorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}
