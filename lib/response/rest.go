package response

import (
	"github.com/gin-gonic/gin"
)

type Response = map[string]interface{}

func Success(data map[string]interface{}, context *gin.Context) {

	response := make(Response)

	response["code"] = 200
	response["msg"] = "ok"
	response["data"] = data

	context.JSON(200, response)
}

func Error(code int, message string, context *gin.Context) {
	response := make(Response)

	response["code"] = code
	response["msg"] = message

	context.JSON(200, response)
}