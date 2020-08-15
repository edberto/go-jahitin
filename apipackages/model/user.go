package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	IUser interface {
		GetAll(param GetAllUserParam) ([]entity.UserEntity, error)
		GetOne(param GetOneUserParam) (entity.UserEntity, error)
		InsertOne(param InsertOneUserParam) (entity.UserEntity, error)
	}

	User struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllUserParam struct {
		IDs []int
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

func (model *User) GetAll(param GetAllUserParam) ([]entity.UserEntity, error) {
	res := new([]entity.UserEntity)

	selectQ := `
		SELECT id, role, address, email, name, username, uuid, tailor_id, created_at, updated_at
		FROM users
	`

	whereQ := ` WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.IDs; len(p) != 0 {
		whereQ += ` AND id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", selectQ, whereQ), 1)

	rows, err := model.Toolkit.DB.Query(q, whereP...)
	if err != nil {
		return *new([]entity.UserEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.UserEntity)

		if err := rows.Scan(&t.ID, &t.Role, &t.Address, &t.Email, &t.Name, &t.Username, &t.UUID, &t.TailorID, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.UserEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
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
	res := new(entity.UserEntity)

	selectQ := `
		SELECT id, role, address, email, name, password, username, uuid, tailor_id, created_at, updated_at
		FROM users
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.ID; p != 0 {
		whereQ += ` AND id = ?`
		whereP = append(whereP, p)
	}
	if p := param.Username; p != "" {
		whereQ += ` AND username = ?`
		whereP = append(whereP, p)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s%s", selectQ, whereQ, limitQ), 1)
	err := model.Toolkit.DB.QueryRow(q, whereP...).Scan(&res.ID, &res.Role, &res.Address, &res.Email, &res.Name, &res.Password, &res.Username, &res.UUID, &res.TailorID, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}
