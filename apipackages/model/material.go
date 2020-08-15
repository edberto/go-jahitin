package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	IMaterial interface {
		GetAll(param GetAllMaterialParam) ([]entity.MaterialEntity, error)
		GetOne(param GetOneMaterialParam) (entity.MaterialEntity, error)
	}

	Material struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllMaterialParam struct {
		IDs []int
	}

	GetOneMaterialParam struct {
		ID int
	}
)

func NewMaterialModel(tk *apipackages.Toolkit) IMaterial {
	return &Material{
		Toolkit: tk,
	}
}

func (model *Material) GetAll(param GetAllMaterialParam) ([]entity.MaterialEntity, error) {
	res := new([]entity.MaterialEntity)

	selectQ := `
		SELECT id, uuid, name, color, detail, created_at, updated_at
		FROM materials
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
		return *new([]entity.MaterialEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.MaterialEntity)

		if err := rows.Scan(&t.ID, &t.UUID, &t.Name, &t.Color, &t.Detail, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.MaterialEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
}

func (model *Material) GetOne(param GetOneMaterialParam) (entity.MaterialEntity, error) {
	res := new(entity.MaterialEntity)

	selectQ := `
		SELECT id, uuid, name, color, detail, created_at, updated_at
		FROM materials
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.ID; p != 0 {
		whereQ += ` AND id = ?`
		whereP = append(whereP, p)
	}

	limitQ := ` ORDER BY updated_at DESC LIMIT 1`

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s%s", selectQ, whereQ, limitQ), 1)

	err := model.Toolkit.DB.QueryRow(q, whereP...).Scan(&res.ID, &res.UUID, &res.Name, &res.Color, &res.Detail, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}
