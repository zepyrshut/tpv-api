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
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (m *Repository) GetAllEnabledItems(c *gin.Context) {
	items, err := m.DB.AllEnabledItems()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func (m *Repository) GetOneItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	item, err := m.DB.OneItem(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}
