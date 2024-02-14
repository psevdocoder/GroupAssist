package rest

import (
	"GroupAssist/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", h.getUsers)
		users.POST("/", h.createUser)
		users.DELETE("/:id", h.deleteUser)
		users.PUT("/:id", h.updateUser)
	}
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) createUser(c *gin.Context) {
	var createUser domain.CreateUser
	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidRequestBody})
		return
	}

	user, err := h.UserService.Create(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}
	err = h.UserService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
	}
	c.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestURL})
		return
	}
	var user domain.UpdateUserInput

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": invalidRequestBody})
		return
	}
	err = h.UserService.Update(id, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": notFound})
	}
	c.AbortWithStatus(http.StatusNoContent)
}
