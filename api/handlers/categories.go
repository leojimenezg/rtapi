package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/services"
)

func (h *Handler) GetAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "categories") })
		return
	}
	c.JSON(http.StatusOK, gin.H{ 
		"records": categories, "total": len(categories) })
}

func (h *Handler) GetCategoryById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString) })
		return
	}
	category, categoryErr := services.GetCategoryById(h.db, idInt)
	if categoryErr != nil {
		if _, ok := categoryErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "category", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "category", idString) })
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) GetCategoryByName(c *gin.Context) {
	name := c.Param("name")
	category, err := services.GetCategoryByName(h.db, name)
	if err != nil {
		if _, ok := err.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NAME_NOT_FOUND, "category", name) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_NAME_ERROR, "category", name) })
		return
	}
	c.JSON(http.StatusOK, category)
}
