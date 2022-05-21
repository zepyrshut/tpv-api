package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Repository) GetAllLounges(c *gin.Context) {

	lounges, err := m.DB.AllLounges()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lounges": lounges,
	})
}
