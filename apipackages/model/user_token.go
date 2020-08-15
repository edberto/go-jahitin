package model

import (
	"fmt"
	"time"

	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"
)

type (
	IUserToken interface {
		GetOne(param GetOneUserTokenParam) (entity.UserTokenEntity, error)
		InsertOne(param InsertOneUserTokenParam) (entity.UserTokenEntity, error)
		DeleteAll(param DeleteAllUserTokenParam) error
	}

	UserToken struct {
		Toolkit *apipackages.Toolkit
	}

	GetOneUserTokenParam struct {
		UserID int
		UUID string
	}

	InsertOneUserTokenParam struct {
		UserID int
		UUID string
		ExpiredAt time.Time
	}

	DeleteAllUserTokenParam struct {
		UserID int
		UUID string
	}
)

func NewUserTokenModel(tk *apipackages.Toolkit) IUserToken {
	return &UserToken{
		Toolkit: tk,
	}
}

func (model *UserToken) GetOne(param GetOneUserTokenParam) (entity.UserTokenEntity, error) {
	res := new(entity.UserTokenEntity)

	selectQ := `
		SELECT user_id, uuid, expired_at
		FROM user_tokens
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.UserID; p != 0 {
		whereQ += ` AND user_id = ?`
		whereP = append(whereP, p)
	}
	if p := param.UUID; p != "" {
		whereQ += ` AND uuid = ?`
		whereP = append(whereP, p)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s%s", selectQ, whereQ, limitQ), 1)
	err := model.Toolkit.DB.QueryRow(q, whereP...).Scan(&res.UserID, &res.UUID, &res.ExpiredAt)

	return *res, err
}

func (model *UserToken) InsertOne(param InsertOneUserTokenParam) (entity.UserTokenEntity, error) {
	res := new(entity.UserTokenEntity)

	q := `
		INSERT INTO user_tokens (user_id, uuid, expired_at)
		VALUES (?, ?, ?)
		RETURNING user_id, uuid, expired_at
	`
	p := []interface{}{param.UserID, param.UUID, param.ExpiredAt}

	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.UserID, &res.UUID, &res.ExpiredAt)

	return *res, err
}

func (model *UserToken) DeleteAll(param DeleteAllUserTokenParam) error {
	deleteQ := `
		DELETE 
		FROM user_tokens
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.UserID; p != 0 {
		whereQ += ` AND user_id = ?`
		whereP = append(whereP, p)
	}
	if p := param.UUID; p != "" {
		whereQ += ` AND uuid = ?`
		whereP = append(whereP, p)
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", deleteQ, whereQ), 1)
	_, err := model.Toolkit.DB.Exec(q, whereP...)

	return err
}