package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/leojimenezg/rtapi/api/services"
)

func (h *Handler) GetAllTopics(c *gin.Context) {
	topics, err := services.GetAllTopics(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "topics") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": topics, "total": len(topics) })
}

func (h *Handler) GetTopicById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	topic, topicErr := services.GetTopicById(h.db, idInt)
	if topicErr != nil {
		if _, ok := topicErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "topic", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "topic", idString) })
		return
	}
	c.JSON(http.StatusOK, topic)
}

func (h *Handler) GetTopicByName(c *gin.Context) {
	name := c.Param("name")
	topic, topicErr := services.GetTopicByName(h.db, name)
	if topicErr != nil {
		if _, ok := topicErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NAME_NOT_FOUND, "topic", name) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_NAME_ERROR, "topic", name) })
		return
	}
	c.JSON(http.StatusOK, topic)
}

func (h *Handler) GetRandomTopic(c *gin.Context) {
	categoryIdString, categoryIdOk := c.GetQuery("category_id")
	difficultyIdString, difficultyIdOk := c.GetQuery("difficulty_id")
	if !categoryIdOk && !difficultyIdOk {
		randomTopic, err := services.GetRandomTopic(h.db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ 
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.ID)
		c.JSON(http.StatusOK, randomTopic)
	} else if categoryIdOk && !difficultyIdOk {
		categoryId, idErr := strconv.ParseInt(categoryIdString, 10, 64)
		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{ 
				"error": fmt.Sprintf(INVALID_ID_FORMAT, categoryIdString) })
			return
		}
		randomTopic, err := services.GetRandomTopicByCategory(h.db, categoryId)
		if err != nil {
			if _, ok := err.(services.NotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{ 
					"error": fmt.Sprintf(NO_TOPICS_FOR_CATEGORY, categoryIdString) })
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{ 
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.ID)
		c.JSON(http.StatusOK, randomTopic)
	} else if difficultyIdOk && !categoryIdOk {
		difficultyId, idErr := strconv.ParseInt(difficultyIdString, 10, 64)
		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{ 
				"error": fmt.Sprintf(INVALID_ID_FORMAT, difficultyIdString) })
			return
		}
		randomTopic, err := services.GetRandomTopicByDifficulty(h.db, difficultyId)
		if err != nil {
			if _, ok := err.(services.NotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{ 
					"error": fmt.Sprintf(NO_TOPICS_FOR_DIFFICULTY, difficultyIdString) })
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{ 
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.ID)
		c.JSON(http.StatusOK, randomTopic)
	} else {
		categoryId, categoryIdErr := strconv.ParseInt(categoryIdString, 10, 64)
		if categoryIdErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf(INVALID_ID_FORMAT, categoryIdString) })
			return
		}
		difficultyId, difficultyIdErr := strconv.ParseInt(difficultyIdString, 10, 64)
		if difficultyIdErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf(INVALID_ID_FORMAT, difficultyIdString) })
			return
		}
		randomTopic, err := services.GetRandomTopicByCategoryAndDifficulty(h.db, categoryId, difficultyId)
		if err != nil {
			if _, ok := err.(services.NotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{ "error": NO_TOPICS_FOR_FILTERS })
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{ 
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.ID)
		c.JSON(http.StatusOK, randomTopic)
	}
}

func (h *Handler) GetAllTopicsWithDetails(c *gin.Context) {
	topicsWithDetails, err := services.GetAllTopicsWithDetails(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_ALL_ERROR, "topics") })
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"records": topicsWithDetails, "total": len(topicsWithDetails) })
}

func (h *Handler) GetTopicWithDetailsById(c *gin.Context) {
	idString := c.Param("id")
	idInt, idErr := strconv.ParseInt(idString, 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf(INVALID_ID_FORMAT, idString)})
		return
	}
	topicWithDetails, topicErr := services.GetTopicWithDetailsById(h.db, idInt)
	if topicErr != nil {
		if _, ok := topicErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(ID_NOT_FOUND, "topic", idString) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_ID_ERROR, "topic", idString) })
		return
	}
	c.JSON(http.StatusOK, topicWithDetails)
}

func (h *Handler) GetTopicWithDetailsByName(c *gin.Context) {
	name := c.Param("name")
	topicWithDetails, topicErr := services.GetTopicsWithDetailsByName(h.db, name)
	if topicErr != nil {
		if _, ok := topicErr.(services.NotFoundError); ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf(NAME_NOT_FOUND, "topic", name) })
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(FETCH_BY_NAME_ERROR, "topic", name) })
		return
	}
	c.JSON(http.StatusOK, topicWithDetails)
}

func (h *Handler) GetRandomTopicWithDetails(c *gin.Context) {
	categoryIdString, categoryIdOk := c.GetQuery("category_id")
	difficultyIdString, difficultyIdOk := c.GetQuery("difficulty_id")
	if !categoryIdOk && !difficultyIdOk {
		randomTopic, err := services.GetRandomTopicWithDetails(h.db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.TopicID)
		c.JSON(http.StatusOK, randomTopic)
	} else if categoryIdOk && !difficultyIdOk {
		categoryId, idErr := strconv.ParseInt(categoryIdString, 10, 64)
		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf(INVALID_ID_FORMAT, categoryIdString) })
			return
		}
		randomTopic, err := services.GetRandomTopicWithDetailsByCategory(h.db, categoryId)
		if err != nil {
			if _, ok := err.(services.NotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf(NO_TOPICS_FOR_CATEGORY, categoryIdString)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.TopicID)
		c.JSON(http.StatusOK, randomTopic)
	} else if difficultyIdOk && !categoryIdOk {
		difficultyId, idErr := strconv.ParseInt(difficultyIdString, 10, 64)
		if idErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf(INVALID_ID_FORMAT, difficultyIdString) })
			return
		}
		randomTopic, err := services.GetRandomTopicWithDetailsByDifficulty(h.db, difficultyId)
		if err != nil {
			if _, ok := err.(services.NotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{
					"error": fmt.Sprintf(NO_TOPICS_FOR_DIFFICULTY, difficultyIdString) })
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.TopicID)
		c.JSON(http.StatusOK, randomTopic)
	} else {
		categoryId, categoryIdErr := strconv.ParseInt(categoryIdString, 10, 64)
		if categoryIdErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf(INVALID_ID_FORMAT, categoryIdString) })
			return
		}
		difficultyId, difficultyIdErr := strconv.ParseInt(difficultyIdString, 10, 64)
		if difficultyIdErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf(INVALID_ID_FORMAT, difficultyIdString) })
			return
		}
		randomTopic, err := services.GetRandomTopicWithDetailsByCategoryAndDifficulty(
			h.db, categoryId, difficultyId)
		if err != nil {
			if _, ok := err.(services.NotFoundError); ok {
				c.JSON(http.StatusNotFound, gin.H{ "error": NO_TOPICS_FOR_FILTERS })
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf(FETCH_ONE_ERROR, "topic") })
			return
		}
		go func(db *sql.DB, id int64) {
			services.UpdateTopicCount(db, id)
		}(h.db, randomTopic.TopicID)
		c.JSON(http.StatusOK, randomTopic)
	}
}
