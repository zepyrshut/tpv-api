package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Repository) GetAllCategories(c *gin.Context) {
	categories, err := m.DB.AllCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed_request",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"categories": categories,
	})
}

func (m *Repository) GetOneCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed_id",
			"error":  err.Error(),
		})
		return
	}

	category, err := m.DB.OneCategory(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not_found",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"category": category,
	})
}
