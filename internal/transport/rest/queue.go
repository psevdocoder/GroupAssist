package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) InitQueuesRoutes(api *gin.RouterGroup) {
	queues := api.Group("/queues")
	{
		queues.GET("/by_subject/:id", h.getAllQueuesBySubject)
		queues.GET("/:id", h.getQueueByID)
		//queues.POST("/", h.createQueue)
		//queues.DELETE("/:id", h.deleteQueue)
		//queues.PUT("/:id", h.updateQueue)
	}
}

func (h *Handler) getAllQueuesBySubject(c *gin.Context) {
	subjectIDStr := c.Param("id")
	subjectID, err := strconv.Atoi(subjectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}

	queues, err := h.QueueService.GetAllBySubject(subjectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
		return
	}
	c.JSON(200, queues)
}

func (h *Handler) getQueueByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}
	queue, err := h.QueueService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
		return
	}
	c.JSON(200, queue)
}
