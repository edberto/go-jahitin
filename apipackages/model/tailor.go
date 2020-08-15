package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	ITailor interface {
		GetAll(param GetAllTailorParam) ([]entity.TailorEntity, error)
		InsertOne(param InsertOneTailorParam) (entity.TailorEntity, error)
		UpdateOne(param UpdateOneTailorParam) (entity.TailorEntity, error)
	}

	Tailor struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllTailorParam struct {
		IDs []int
	}

	InsertOneTailorParam struct {
		Address string
		Email   string
		Name    string
		UUID    string
	}

	UpdateOneTailorParam struct{}
)

func NewTailorModel(tk *apipackages.Toolkit) ITailor {
	return &Tailor{
		Toolkit: tk,
	}
}

func (model *Tailor) GetAll(param GetAllTailorParam) ([]entity.TailorEntity, error) {
	res := new([]entity.TailorEntity)

	selectQ := `
		SELECT id, uuid, name, address, email, created_at, updated_at
		FROM tailors
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.IDs; len(p) != 0 {
		whereQ += ` AND id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", selectQ, whereQ), 1)

	rows, err := model.Toolkit.DB.Query(q, whereP...)
	if err != nil {
		return *new([]entity.TailorEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.TailorEntity)

		if err := rows.Scan(&t.ID, &t.UUID, &t.Name, &t.Address, &t.Email, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.TailorEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
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

func (model *Tailor) UpdateOne(param UpdateOneTailorParam) (entity.TailorEntity, error) {
	res := new(entity.TailorEntity)

	return *res, nil
}
