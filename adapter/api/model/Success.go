package model

import (
	"github.com/gin-gonic/gin"
)

type Success struct {
	statusCode int
	result     interface{}
}

func NewSuccess(result interface{}, status int) Success {
	return Success{
		statusCode: status,
		result:     result,
	}
}

func (r Success) Send(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(r.statusCode, r.result)
}
