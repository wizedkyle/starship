package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryLimit(c *gin.Context) int64 {
	limit := c.DefaultQuery("limit", "10")
	limitInt64, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0
	}
	return limitInt64
}
