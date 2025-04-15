package message

import (
	"context"

	"github.com/Khmelov/Distcomp/251003/Nasevich/restbasics/discussion/internal/model"
	"github.com/gin-gonic/gin"
)

type messageService interface {
	GetMessage(ctx context.Context, id int64) (model.Message, error)
	GetMessages(ctx context.Context) ([]model.Message, error)
	CreateMessage(ctx context.Context, args model.Message) (model.Message, error)
	UpdateMessage(ctx context.Context, args model.Message) (model.Message, error)
	DeleteMessage(ctx context.Context, id int64) error
}

type noticeHandler struct {
	notice messageService
}

func New(noticeSvc messageService) *noticeHandler {
	return &noticeHandler{
		notice: noticeSvc,
	}
}

func (h *noticeHandler) InitRoutes(router gin.IRouter) {
	v1 := router.Group("/v1.0")
	{
		v1.GET("/messages", h.List())
		v1.GET("/messages/:id", h.Get())
		v1.POST("/messages", h.Create())
		v1.DELETE("/messages/:id", h.Delete())
		v1.PUT("/messages", h.Update())
	}
}
