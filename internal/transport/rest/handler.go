package rest

import (
	"GroupAssist/internal/domain"
	"GroupAssist/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	notFound           = "Not found"
	invalidRequestURL  = "Invalid request url"
	invalidRequestBody = "Invalid request body"
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

type Handler struct {
	SubjectService Subject
	QueueService   Queue
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		SubjectService: services.Subject,
		QueueService:   services.Queue,
	}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("ping/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func (h *Handler) InitAPI(r *gin.Engine) {
	api := r.Group("/api")

	h.InitSubjectsRoutes(api)
	h.InitQueuesRoutes(api)

}
