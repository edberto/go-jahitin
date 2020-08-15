package handler

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/usecase"

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
		Toolkit   *apipackages.Toolkit
	}
)

func NewSessionHandler(tk *apipackages.Toolkit) ISession {
	return &Session{
		SessionUC: usecase.NewSessionUC(tk),
		Toolkit:   tk,
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
