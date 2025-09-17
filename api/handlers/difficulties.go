package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/services"
)

func (h *Handler) GetAllDifficulties(c *gin.Context) {
	difficulties, err := services.GetAllDifficulties(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "difficulties") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": difficulties, "total": len(difficulties) })
}

func (h *Handler) GetDifficultyById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	difficulty, difficultyErr := services.GetDifficultyById(h.db, idInt)
	if difficultyErr != nil {
		if _, ok := difficultyErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "difficulty", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "difficulty", idString) })
		return
	}
	c.JSON(http.StatusOK, difficulty)
}

func (h *Handler) GetDifficultyByName(c *gin.Context) {
	name := c.Param("name")
	difficulty, err := services.GetDifficultyByName(h.db, name)
	if err != nil {
		if _, ok := err.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NAME_NOT_FOUND, "difficulty", name) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_NAME_ERROR, "difficulty", name) })
		return
	}
	c.JSON(http.StatusOK, difficulty)
}
