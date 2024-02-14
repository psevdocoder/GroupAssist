package rest

import (
	_ "GroupAssist/docs"
	"GroupAssist/internal/domain"
	"GroupAssist/internal/service"
	"GroupAssist/pkg/pkg-gin"
	"github.com/gin-gonic/gin"
	cache "github.com/psevdocoder/InMemoryCacheTTL"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // pkg-gin-swagger middleware
	"time"
)

const (
	notFound                      = "Not found"
	invalidRequestURL             = "Invalid request url"
	invalidRequestBody            = "Invalid request body"
	providedCredentialsAreInvalid = "Provided credentials are invalid"
	AccountAlreadyExists          = "Account with provided username already exists"
)

type Subject interface {
	GetAll() ([]domain.Subject, error)
	GetByID(id int) (domain.Subject, error)
	Create(subject domain.Subject) (domain.Subject, error)
	Delete(id int) error
	Update(id int, subject domain.UpdateSubjectInput) error
}

type Queue interface {
	GetAllBySubject(id int) ([]domain.QueuesResponse, error)
	GetByID(id int) (domain.QueueResponse, error)
	Create(queue domain.Queue) (domain.Queue, error)
	Delete(id int) error
	Update(id int, queue domain.UpdateQueueInput) error
}

type Auth interface {
	ApplyRegister(user domain.SignUpInput) (domain.ResponseUser, error)
	CreateToken(input domain.SignInInput, ip string) (domain.SignInResponse, error)
	RefreshToken(token string) (domain.SignInResponse, error)
}

type User interface {
	Create(user domain.CreateUser) (domain.CreateUser, error)
	Update(id int, user domain.UpdateUserInput) error
	Delete(id int) error
	GetAll() ([]domain.User, error)
}

type Handler struct {
	SubjectService Subject
	QueueService   Queue
	AuthService    Auth
	UserService    User
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		SubjectService: services.Subject,
		QueueService:   services.Queue,
		AuthService:    services.Auth,
		UserService:    services.User,
	}
}

func (h *Handler) Init(memoryCache *cache.Cache, cacheTTL time.Duration) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), pkg_gin.CacheMiddleware(memoryCache, cacheTTL))

	r.GET("ping/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	h.initAuthRoutes(api)
	h.initSubjectsRoutes(api)
	h.initQueuesRoutes(api)
	h.initUsersRoutes(api)

	return r
}
