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

func (h *Handler) getAllSubjects(c *gin.Context) {
	subjects, err := h.SubjectService.GetAll()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, subjects)
}

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
