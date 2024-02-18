package rest

import (
	"GroupAssist/internal/domain"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/jwt/create", h.createToken)
		auth.POST("/jwt/refresh", h.refreshToken)
	}
}

// @Summary Sign up
// @ID sign-up
// @Produce json
// @Description Sign up. Provide there login, password and registration token
// @Param input body domain.SignUpInput true "Credentials"
// @Success 200 {object} domain.ResponseUser
// @Failure 400
// @Failure 401
// @Failure 500
// @Tags auth
// @Router /api/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var signUpInput domain.SignUpInput
	if err := c.ShouldBindJSON(&signUpInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": providedCredentialsAreInvalid})
		return
	}

	user, err := h.AuthService.ApplyRegister(signUpInput)
	if err != nil {
		var pqErr *pq.Error
		switch {
		case errors.Is(err, sql.ErrNoRows):
			c.JSON(http.StatusUnauthorized, gin.H{"error": providedCredentialsAreInvalid})
		case errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation":
			c.JSON(http.StatusConflict, gin.H{"error": AccountAlreadyExists})
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Create JWT
// @ID create-jwt
// @Produce json
// @Param input body domain.SignInInput true "Credentials"
// @Success 200 {object} domain.SignInResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Tags auth
// @Router /api/auth/jwt/create [post]
func (h *Handler) createToken(c *gin.Context) {
	var signInInput domain.SignInInput
	if err := c.ShouldBindJSON(&signInInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidRequestBody})
		return
	}

	ip := c.ClientIP()

	token, err := h.AuthService.CreateToken(signInInput, ip)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": providedCredentialsAreInvalid})
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.JSON(http.StatusOK, token)
}

// @Summary Refresh JWT
// @ID refresh-jwt
// @Produce json
// @Param input body domain.RefreshTokenInput true "Credentials"
// @Success 200 {object} domain.SignInResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Tags auth
// @Router /api/auth/jwt/refresh [post]
func (h *Handler) refreshToken(c *gin.Context) {
	var request domain.RefreshTokenInput

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": invalidRequestBody})
		return
	}

	ip := c.ClientIP()

	token, err := h.AuthService.RefreshToken(request.RefreshToken, ip)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRefreshTokenExpired):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrInvalidSignature):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrInvalidToken):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrUnhandledToken):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, domain.ErrUnexpectedSigningToken):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, sql.ErrNoRows):
			c.JSON(http.StatusUnauthorized, gin.H{"error": domain.ErrInvalidToken.Error()})
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.JSON(http.StatusOK, token)
}
