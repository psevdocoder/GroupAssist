package rest

import (
	"GroupAssist/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/jwt/create", h.generateToken)
		auth.POST("/jwt/refresh", h.refreshToken)
	}
}

func (h *Handler) generateToken(c *gin.Context) {
	var signInInput domain.SignInInput
	if err := c.ShouldBindJSON(&signInInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidRequestBody})
		return
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var signUpInput domain.SignUpInput
	if err := c.ShouldBindJSON(&signUpInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidRequestBody})
		return
	}

	user, err := h.AuthService.ApplyRegister(signUpInput)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) refreshToken(c *gin.Context) {

}
