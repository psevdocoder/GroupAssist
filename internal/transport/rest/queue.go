package rest

import (
	"GroupAssist/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initQueuesRoutes(api *gin.RouterGroup) {
	queues := api.Group("/queues")
	{
		queues.GET("/by_subject/:id", h.getAllQueuesBySubject)
		queues.GET("/:id", h.getQueueByID)
		queues.POST("/", h.createQueue)
		queues.DELETE("/:id", h.deleteQueue)
		queues.PATCH("/:id", h.updateQueue)
		queues.POST("/join", h.joinQueue)
		queues.DELETE("/leave:id", h.leaveQueue)
	}
}

// @Summary Get all queues by subject id
// @ID get-all-queues
// @Produce json
// @Param id path int true "Subject ID"
// @Success 200 {object} []domain.Queue
// @Failure 400
// @Failure 500
// @Tags queues
// @Router /api/queues/by_subject/{id} [get]
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
	c.JSON(http.StatusOK, queues)
}

// @Summary Get queue by ID
// @ID get-queue-by-id
// @Produce json
// @Param id path int true "Queue ID"
// @Success 200 {object} domain.Queue
// @Failure 400
// @Failure 404
// @Tags queues
// @Router /api/queues/{id} [get]
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
	c.JSON(http.StatusOK, queue)
}

// @Summary Create queue
// @ID create-queue
// @Produce json
// @Param input body domain.Queue true "Queue data"
// @Success 200 {object} domain.Queue
// @Failure 400
// @Failure 500
// @Tags queues
// @Router /api/queues [post]
func (h *Handler) createQueue(c *gin.Context) {
	var input domain.Queue
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestBody})
		return
	}
	queue, err := h.QueueService.Create(input)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, queue)
}

// @Summary Delete queue
// @ID delete-queue
// @Produce json
// @Param id path int true "Queue ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Tags queues
// @Router /api/queues/{id} [delete]
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

// @Summary Update queue
// @ID update-queue
// @Produce json
// @Description Allowed to use any field provided in the input body
// @Param input body domain.UpdateQueueInput true "Queue"
// @Param id path int true "Queue ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Tags queues
// @Router /api/queues/{id} [patch]
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
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) joinQueue(context *gin.Context) {
}

func (h *Handler) leaveQueue(context *gin.Context) {
}
