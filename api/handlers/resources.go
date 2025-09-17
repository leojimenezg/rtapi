package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/services"
)

func (h *Handler) GetAllResources(c *gin.Context) {
	resources, err := services.GetAllResources(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resources") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resources, "total": len(resources) })
}

func (h *Handler) GetResourceById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resource, resourceErr := services.GetResourceById(h.db, idInt)
	if resourceErr != nil {
		if _, ok := resourceErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "resource", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "resource", idString) })
		return
	}
	c.JSON(http.StatusOK, resource)
}

func (h *Handler) GetResourcesByType(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resources, resourcesErr := services.GetResourcesByType(h.db, idInt)
	if resourcesErr != nil {
		if _, ok := resourcesErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NO_RESOURCES_FOR_TYPE, idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resources") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resources, "total": len(resources) })
}

func (h *Handler) GetResourcesByTopic(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resources, resourcesErr := services.GetResourcesByTopic(h.db, idInt)
	if resourcesErr != nil {
		if _, ok := resourcesErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NO_RESOURCES_FOR_TOPIC, idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resources") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resources, "total": len(resources) })
}

func (h *Handler) GetAllResourcesWithDetails(c *gin.Context) {
	resourcesWithDetails, err := services.GetAllResourcesWithDetails(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resources") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resourcesWithDetails, "total": len(resourcesWithDetails) })
}

func (h *Handler) GetResourceWithDetailsById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resourceWithDetails, resourceErr := services.GetResourceWithDetailsById(h.db, idInt)
	if resourceErr != nil {
		if _, ok := resourceErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "resource", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "resource", idString) })
		return
	}
	c.JSON(http.StatusOK, resourceWithDetails)
}

func (h *Handler) GetResourcesWithDetailsByType(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resourcesWithDetails, resourcesErr := services.GetResourcesWithDetailsByType(h.db, idInt)
	if resourcesErr != nil {
		if _, ok := resourcesErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NO_RESOURCES_FOR_TYPE, idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resources") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resourcesWithDetails, "total": len(resourcesWithDetails) })
}

func (h *Handler) GetResourcesWithDetailsByTopic(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	resourcesWithDetails, resourcesErr := services.GetResourcesWithDetailsByTopic(h.db, idInt)
	if resourcesErr != nil {
		if _, ok := resourcesErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NO_RESOURCES_FOR_TOPIC, idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "resources") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": resourcesWithDetails, "total": len(resourcesWithDetails) })
}
