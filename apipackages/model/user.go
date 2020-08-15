package model

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"
)

type (
	IUser interface {
		InsertOne(param InsertOneUserParam) (entity.UserEntity, error)
		GetOne(param GetOneUserParam) (entity.UserEntity, error)
	}

	User struct {
		Toolkit *apipackages.Toolkit
	}

	InsertOneUserParam struct {
		Role     int
		TailorID int
		Address  string
		Email    string
		Name     string
		Password string
		Username string
		UUID     string
	}

	GetOneUserParam struct {
		ID       int
		Username string
	}
)

func NewUserModel(tk *apipackages.Toolkit) IUser {
	return &User{
		Toolkit: tk,
	}
}

func (model *User) InsertOne(param InsertOneUserParam) (entity.UserEntity, error) {
	res := new(entity.UserEntity)

	q := `
		INSERT INTO users (role, tailor_id, address, email, name, password, username, uuid) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?) 
		RETURNING id, role, address, email, name, password, username, uuid, tailor_id, created_at, updated_at
	`
	p := []interface{}{param.Role, param.TailorID, param.Address, param.Email, param.Name, param.Password, param.Username, param.UUID}
	
	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.ID, &res.Role, &res.Address, &res.Email, &res.Name, &res.Password, &res.Username, &res.UUID, &res.TailorID, &res.CreatedAt, &res.UpdatedAt)
	
	return *res, err
}

func (model *User) GetOne(param GetOneUserParam) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}
