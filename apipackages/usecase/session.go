package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
)

type (
	ISession interface {
		Login(param LoginSessionParam) (viewmodel.LoginVM, error)
		Logout(param LogoutSessionParam) (viewmodel.LogoutVM, error)
		Refresh(param RefreshSessionParam) (viewmodel.RefreshVM, error)
	}

	Session struct {
		SessionModel model.ISession
		Toolkit      *apipackages.Toolkit
	}

	LoginSessionParam   struct{}
	LogoutSessionParam  struct{}
	RefreshSessionParam struct{}
)

func NewSessionUC(tk *apipackages.Toolkit) ISession {
	return &Session{
		SessionModel: model.NewSessionModel(tk),
		Toolkit:      tk,
	}
}

func (uc *Session) Login(param LoginSessionParam) (viewmodel.LoginVM, error) {
	return viewmodel.LoginVM{}, nil
}

func (uc *Session) Logout(param LogoutSessionParam) (viewmodel.LogoutVM, error) {
	return viewmodel.LogoutVM{}, nil
}

func (uc *Session) Refresh(param RefreshSessionParam) (viewmodel.RefreshVM, error) {
	return viewmodel.RefreshVM{}, nil
}
