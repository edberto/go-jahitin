package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	ITailorMaterial interface {
		GetAll(param GetAllTailorMaterialParam) ([]entity.TailorMaterialEntity, error)
		InsertBulk(param []InsertBulkTailorMaterialParam) (int, error)
	}

	TailorMaterial struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllTailorMaterialParam struct {
		TailorIDs   []int
		MaterialIDs []int
	}

	InsertBulkTailorMaterialParam struct {
		TailorID   int
		MaterialID int
		Price      float64
	}
)

func NewTailorMaterialModel(tk *apipackages.Toolkit) ITailorMaterial {
	return &TailorMaterial{
		Toolkit: tk,
	}
}

func (model *TailorMaterial) GetAll(param GetAllTailorMaterialParam) ([]entity.TailorMaterialEntity, error) {
	res := new([]entity.TailorMaterialEntity)

	selectQ := `
		SELECT id, tailor_id, material_id, price, created_at, updated_at
		FROM tailor_materials
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.TailorIDs; len(p) != 0 {
		whereQ += ` AND tailor_id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}
	if p := param.MaterialIDs; len(p) != 0 {
		whereQ += ` AND material_id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", selectQ, whereQ), 1)

	rows, err := model.Toolkit.DB.Query(q, whereP...)
	if err != nil {
		return *new([]entity.TailorMaterialEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.TailorMaterialEntity)

		if err := rows.Scan(&t.ID, &t.TailorID, &t.MaterialID, &t.Price, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.TailorMaterialEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
}

func (model *TailorMaterial) InsertBulk(param []InsertBulkTailorMaterialParam) (int, error) {
	if len(param) <= 0 {
		return 0, nil
	}

	trx, err := model.Toolkit.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer trx.Rollback()

	q := `INSERT INTO tailor_materials (tailor_id, material_id, price) VALUES`
	p := []interface{}{}
	for i, v := range param {
		q += ` (?, ?, ?)`
		p = append(p, v.TailorID, v.MaterialID, v.Price)

		if i != len(param)-1 {
			q += `,`
		}
	}

	q = helper.ReplacePlaceholder(q, 1)

	stmt, err := trx.Prepare(q)
	if err != nil {
		if err := trx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	result, err := stmt.Exec(p...)
	if err != nil {
		if err := trx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		if err := trx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	if err := trx.Commit(); err != nil {
		if err := trx.Rollback(); err != nil {
			return 0, err
		}
		return 0, err
	}

	return int(rows), nil
}
