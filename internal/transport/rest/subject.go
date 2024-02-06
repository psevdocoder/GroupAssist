package rest

import (
	"GroupAssist/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) InitSubjectsRoutes(api *gin.RouterGroup) {
	subjects := api.Group("/subjects")
	{
		subjects.GET("/", h.getAllSubjects)
		subjects.GET("/:id", h.getSubjectByID)
		subjects.POST("/", h.createSubject)
		subjects.DELETE("/:id", h.deleteSubject)
		subjects.PUT("/:id", h.updateSubject)
	}
}

// @Summary Get all subjects
// @ID get-all-subjects
// @Produce json
// @Success 200 {object} []domain.Subject
// @Failure 500
// @Tags subjects
// @Router /api/subjects [get]
func (h *Handler) getAllSubjects(c *gin.Context) {
	subjects, err := h.SubjectService.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, subjects)
}

// @Summary Get subject by ID
// @ID get-subject-by-id
// @Produce json
// @Param id path int true "Subject ID"
// @Success 200 {object} domain.Subject
// @Failure 400
// @Failure 404
// @Tags subjects
// @Router /api/subjects/{id} [get]
func (h *Handler) getSubjectByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}

	subject, err := h.SubjectService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
		return
	}
	c.JSON(200, subject)
}

// @Summary Create subject
// @ID create-subject
// @Produce json
// @Param input body domain.Subject true "Subject data"
// @Success 200 {object} domain.Subject
// @Failure 400
// @Failure 500
// @Tags subjects
// @Router /api/subjects [post]
func (h *Handler) createSubject(c *gin.Context) {
	var subject domain.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestBody})
		return
	}
	createdSubject, err := h.SubjectService.Create(subject)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(200, createdSubject)
}

// @Summary Delete subject
// @ID delete-subject
// @Produce json
// @Param id path int true "Subject ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Tags subjects
// @Router /api/subjects{id} [delete]
func (h *Handler) deleteSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}
	err = h.SubjectService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

// @Summary Update subject
// @ID update-subject
// @Produce json
// @Param id path int true "Subject ID"
// @Param input body domain.UpdateSubjectInput true "Subject data"
// @Success 204
// @Failure 400
// @Failure 500
// @Tags subjects
// @Router /api/subjects/{id} [put]
func (h *Handler) updateSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}
	var subject domain.UpdateSubjectInput
	if err = c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestBody})
		return
	}
	err = h.SubjectService.Update(id, subject)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}
