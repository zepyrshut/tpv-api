package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Repository) GetTableFromLounge(c *gin.Context) {

	loungeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
		})
		return
	}

	tables, err := m.DB.AllTablesFromSelectedLounge(loungeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
