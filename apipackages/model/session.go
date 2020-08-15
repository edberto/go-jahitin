package model

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
)

type (
	ISession interface {
		InsertOne(param InsertOneSessionParam) (entity.SessionEntity, error)
		DeleteOne(param DeleteOneSessionParam) error
	}

	Session struct {
		Toolkit *apipackages.Toolkit
	}

	InsertOneSessionParam struct{}
	DeleteOneSessionParam struct{}
)

func NewSessionModel(tk *apipackages.Toolkit) ISession {
	return &Session{
		Toolkit: tk,
	}
}

func (model *Session) InsertOne(param InsertOneSessionParam) (entity.SessionEntity, error) {
	return entity.SessionEntity{}, nil
}

func (model *Session) DeleteOne(param DeleteOneSessionParam) error {
	return nil
}
