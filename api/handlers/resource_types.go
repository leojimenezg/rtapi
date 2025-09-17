package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/services"
)

func (h *Handler) GetAllResourceTypes(c *gin.Context) {
	resourceTypes, err := services.GetAllResourceTypes(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resource_types") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resourceTypes, "total": len(resourceTypes) })
}

func (h *Handler) GetResourceTypeById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resourceType, resourceTypeErr := services.GetResourceTypeById(h.db, idInt)
	if resourceTypeErr != nil {
		if _, ok := resourceTypeErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "resource_type", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "resource_type", idString) })
		return
	}
	c.JSON(http.StatusOK, resourceType)
}

func (h *Handler) GetResourceTypeByName(c *gin.Context) {
	name := c.Param("name")
	resourceType, err := services.GetResourceTypeByName(h.db, name)
	if err != nil {
		if _, ok := err.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NAME_NOT_FOUND, "resource_type", name) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_NAME_ERROR, "resource_type", name) })
		return
	}
	c.JSON(http.StatusOK, resourceType)
}
