package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	IModel interface {
		GetAll(param GetAllModelParam) ([]entity.ModelEntity, error)
		GetOne(param GetOneModelParam) (entity.ModelEntity, error)
	}

	Model struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllModelParam struct {
		IDs []int
	}

	GetOneModelParam struct {
		ID int
	}
)

func NewModelModel(tk *apipackages.Toolkit) IModel {
	return &Model{
		Toolkit: tk,
	}
}

func (model *Model) GetAll(param GetAllModelParam) ([]entity.ModelEntity, error) {
	res := new([]entity.ModelEntity)

	selectQ := `
		SELECT id, uuid, name, detail, created_at, updated_at
		FROM models
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
		return *new([]entity.ModelEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.ModelEntity)

		if err := rows.Scan(&t.ID, &t.UUID, &t.Name, &t.Detail, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.ModelEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
}

func (model *Model) GetOne(param GetOneModelParam) (entity.ModelEntity, error) {
	res := new(entity.ModelEntity)

	selectQ := `
		SELECT id, uuid, name, detail, created_at, updated_at
		FROM models
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.ID; p != 0 {
		whereQ += ` AND id = ?`
		whereP = append(whereP, p)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s%s", selectQ, whereQ, limitQ), 1)

	err := model.Toolkit.DB.QueryRow(q, whereP...).Scan(&res.ID, &res.UUID, &res.Name, &res.Detail, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}
