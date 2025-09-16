package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) GetAllTopics(c *gin.Context) {}

func (h *Handler) GetTopicById(c *gin.Context) {}

func (h *Handler) GetTopicByName(c *gin.Context) {}

func (h *Handler) GetAllTopicsWithDetails(c *gin.Context) {}

func (h *Handler) GetTopicWithDetailsById(c *gin.Context) {}

func (h *Handler) GetTopicWithDetailsByName(c *gin.Context) {}

func (h *Handler) GetRandomTopicWithDetails(c *gin.Context) {}

func (h *Handler) GetRandomTopicWithDetailsByCategory(c *gin.Context) {}

func (h *Handler) GetRandomTopicWithDetailsByDifficulty(c *gin.Context) {}

func (h *Handler) GetRandomTopicWithDetailsByCategoryAndDifficulty(c *gin.Context) {}
