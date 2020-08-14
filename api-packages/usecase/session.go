package usecase

import (
	"go-jahitin/api-packages/model"
	"go-jahitin/api-packages/viewmodel"

	"github.com/gin-gonic/gin"
)

type (
	ISession interface {
		Login(c *gin.Context, param LoginSessionParam) (viewmodel.LoginVM, error)
		Logout(c *gin.Context, param LogoutSessionParam) (viewmodel.LogoutVM, error)
		Refresh(c *gin.Context, param RefreshSessionParam) (viewmodel.RefreshVM, error)
	}

	Session struct {
		SessionModel model.ISession
	}

	LoginSessionParam   struct{}
	LogoutSessionParam  struct{}
	RefreshSessionParam struct{}
)

func NewSessionUC() ISession {
	return &Session{
		SessionModel: model.NewSessionModel(),
	}
}

func (uc *Session) Login(c *gin.Context, param LoginSessionParam) (viewmodel.LoginVM, error) {
	return viewmodel.LoginVM{}, nil
}

func (uc *Session) Logout(c *gin.Context, param LogoutSessionParam) (viewmodel.LogoutVM, error) {
	return viewmodel.LogoutVM{}, nil
}

func (uc *Session) Refresh(c *gin.Context, param RefreshSessionParam) (viewmodel.RefreshVM, error) {
	return viewmodel.RefreshVM{}, nil
}
