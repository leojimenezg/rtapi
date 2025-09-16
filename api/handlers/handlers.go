package handlers

import (
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/services"
)

type Handler struct {
	db *sql.DB
}

func New(database *sql.DB) *Handler {
	return &Handler{ db: database }
}

func (h *Handler) GetAllTopics(c *gin.Context) {
	topics, err := services.GetAllTopics(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": "failed to fetch all topics" })
	}
	c.JSON(http.StatusOK, gin.H{ "records": topics })
}

func (h *Handler) GetAllTopicsWithDetails(c *gin.Context) {
	topics, err := services.GetAllTopicsWithDetails(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "error": "failed to fetch all topics with details" })
	}
	c.JSON(http.StatusOK, gin.H{ "records": topics })
}
