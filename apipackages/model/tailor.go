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
		GetOne(param GetOneTailorParam) (entity.TailorEntity, error)
		InsertOne(param InsertOneTailorParam) (entity.TailorEntity, error)
		UpdateOne(param UpdateOneTailorParam) (entity.TailorEntity, error)
	}

	Tailor struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllTailorParam struct {
		IDs     []int
		Keyword string
	}

	GetOneTailorParam struct {
		ID int
	}

	InsertOneTailorParam struct {
		Phone string
		Email string
		UUID  string
	}

	UpdateOneTailorParam struct {
		Name    string
		Address string
		ID      int
	}
)

func NewTailorModel(tk *apipackages.Toolkit) ITailor {
	return &Tailor{
		Toolkit: tk,
	}
}

func (model *Tailor) GetAll(param GetAllTailorParam) ([]entity.TailorEntity, error) {
	res := new([]entity.TailorEntity)

	selectQ := `
		SELECT id, uuid, name, address, email, phone, created_at, updated_at
		FROM tailors
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.IDs; len(p) != 0 {
		whereQ += ` AND id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}
	if p := param.Keyword; p != "" {
		whereQ += ` AND name ILIKE ?`
		whereP = append(whereP, fmt.Sprintf("%%%s%%", p))
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", selectQ, whereQ), 1)

	rows, err := model.Toolkit.DB.Query(q, whereP...)
	if err != nil {
		return *new([]entity.TailorEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.TailorEntity)

		if err := rows.Scan(&t.ID, &t.UUID, &t.Name, &t.Address, &t.Email, &t.Phone, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.TailorEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
}

func (model *Tailor) GetOne(param GetOneTailorParam) (entity.TailorEntity, error) {
	res := new(entity.TailorEntity)

	selectQ := `
	SELECT id, uuid, name, address, email, phone, created_at, updated_at
	FROM tailors
`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.ID; p != 0 {
		whereQ += ` AND id = ?`
		whereP = append(whereP, p)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s%s", selectQ, whereQ, limitQ), 1)

	err := model.Toolkit.DB.QueryRow(q, whereP...).Scan(&res.ID, &res.UUID, &res.Name, &res.Address, &res.Email, &res.Phone, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}

func (model *Tailor) InsertOne(param InsertOneTailorParam) (entity.TailorEntity, error) {
	res := new(entity.TailorEntity)

	q := `
		INSERT INTO tailors (phone, email, uuid)
		VALUES (?, ?, ?)
		RETURNING id, phone, address, email, name, uuid, created_at, updated_at
	`
	p := []interface{}{param.Phone, param.Email, param.UUID}

	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.ID, &res.Phone, &res.Address, &res.Email, &res.Name, &res.UUID, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}

func (model *Tailor) UpdateOne(param UpdateOneTailorParam) (entity.TailorEntity, error) {
	res := new(entity.TailorEntity)

	q := `
		UPDATE tailors
		SET updated_at = NOW(), name = ?, address = ?
		WHERE id = ?
		RETURNING id, phone, address, email, name, uuid, created_at, updated_at
	`
	p := []interface{}{param.Name, param.Address, param.ID}

	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.ID, &res.Phone, &res.Address, &res.Email, &res.Name, &res.UUID, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}
