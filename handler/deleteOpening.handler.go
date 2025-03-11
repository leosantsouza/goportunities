package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpening(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "DELETE Opening"})
}
