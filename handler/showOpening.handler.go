package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "GET Opening"})
}
