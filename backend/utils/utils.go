package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func WriteJSON(c *gin.Context, status int, v any) {
	c.JSON(status, v)
}

func WriteError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

func ParseJSON(c *gin.Context, v any) error {
	if c.Request.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(c.Request.Body).Decode(v)
}

func GetTokenFromRequest(c *gin.Context) string {
	tokenAuth := c.GetHeader("Authorization")
	tokenQuery := c.Query("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
