package model

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"
)

type (
	ITailor interface {
		InsertOne(param InsertOneTailorParam) (entity.TailorEntity, error)
	}

	Tailor struct {
		Toolkit *apipackages.Toolkit
	}

	InsertOneTailorParam struct {
		Address string
		Email   string
		Name    string
		UUID    string
	}
)

func NewTailorModel(tk *apipackages.Toolkit) ITailor {
	return &Tailor{
		Toolkit: tk,
	}
}

func (model *Tailor) InsertOne(param InsertOneTailorParam) (entity.TailorEntity, error) {
	res := new(entity.TailorEntity)

	q := `
		INSERT INTO tailors (address, email, name, uuid)
		VALUES (?, ?, ?, ?)
		RETURNING id, address, email, name, uuid, created_at, updated_at
	`
	p := []interface{}{param.Address, param.Email, param.Name, param.UUID}

	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.ID, &res.Address, &res.Email, &res.Name, &res.UUID, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}
