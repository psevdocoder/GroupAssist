package rest

import (
	"GroupAssist/internal/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) InitQueuesRoutes(api *gin.RouterGroup) {
	queues := api.Group("/queues")
	{
		queues.GET("/by_subject/:id", h.getAllQueuesBySubject)
		queues.GET("/:id", h.getQueueByID)
		queues.POST("/", h.createQueue)
		queues.DELETE("/:id", h.deleteQueue)
		queues.PATCH("/:id", h.updateQueue)
		queues.POST("/join", h.joinQueue)
		queues.DELETE("/leave", h.leaveQueue)
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
		c.AbortWithStatus(http.StatusInternalServerError)
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

func (h *Handler) createQueue(c *gin.Context) {
	var input domain.Queue
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestBody})
		return
	}
	queue, err := h.QueueService.Create(input)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, queue)
}

func (h *Handler) deleteQueue(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}
	err = h.QueueService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) updateQueue(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}

	_, err = h.QueueService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
		return
	}

	var input domain.UpdateQueueInput
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestBody})
		return
	}

	err = h.QueueService.Update(id, input)
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) joinQueue(context *gin.Context) {
	//TODO
}

func (h *Handler) leaveQueue(context *gin.Context) {
	//TODO
}
