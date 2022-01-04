package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func badRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
}

func parseFloatQuery(c *gin.Context, key string) (*float64, error) {
	stringVal, present := c.GetQuery(key)
	if !present {
		return nil, nil
	}
	float, err := strconv.ParseFloat(stringVal, 64)
	if err != nil {
		return nil, err
	}
	return &float, nil
}
