package http

import "github.com/gin-gonic/gin"

func ParseRequest(context *gin.Context, data interface{}) error {
	err := context.ShouldBindJSON(&data)
	if err != nil {
		print("error while parsing the data")
	}

	return nil
}
