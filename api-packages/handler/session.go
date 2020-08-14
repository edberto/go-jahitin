package handler

import (
	"go-jahitin/api-packages/usecase"

	"github.com/gin-gonic/gin"
)

type (
	ISession interface {
		Login(c *gin.Context)
		Logout(c *gin.Context)
		Refresh(c *gin.Context)
	}

	Session struct {
		SessionUC usecase.ISession
	}
)

func NewSessionHandler() ISession {
	return &Session{
		SessionUC: usecase.NewSessionUC(),
	}
}

func (h *Session) Login(c *gin.Context) {
	return
}

func (h *Session) Logout(c *gin.Context) {
	return
}

func (h *Session) Refresh(c *gin.Context) {
	return
}
