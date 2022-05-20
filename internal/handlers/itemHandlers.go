package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (m *Repository) GetAllItems(c *gin.Context) {
	items, err := m.DB.AllItems()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "not_found",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"items":  items,
	})
}

func (m *Repository) GetAllEnabledItems(c *gin.Context) {
	items, err := m.DB.AllEnabledItems()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not_found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"items":  items,
	})
}

func (m *Repository) GetOneItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed_id",
			"error":  err.Error(),
		})
		return
	}

	item, err := m.DB.OneItem(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not_found",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"item":   item,
	})
}
