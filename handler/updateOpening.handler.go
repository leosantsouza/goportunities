package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "PUT Opening"})
}
