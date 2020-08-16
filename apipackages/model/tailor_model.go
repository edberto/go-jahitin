package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	ITailorModel interface {
		GetAll(param GetAllTailorModelParam) ([]entity.TailorModelEntity, error)
		InsertBulk(param []InsertBulkTailorModelParam) (int, error)
	}

	TailorModel struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllTailorModelParam struct {
		TailorIDs []int
		ModelIDs  []int
	}

	InsertBulkTailorModelParam struct {
		TailorID int
		ModelID  int
		Price    float64
	}
)

func NewTailorModelModel(tk *apipackages.Toolkit) ITailorModel {
	return &TailorModel{
		Toolkit: tk,
	}
}

func (model *TailorModel) GetAll(param GetAllTailorModelParam) ([]entity.TailorModelEntity, error) {
	res := new([]entity.TailorModelEntity)

	selectQ := `
		SELECT id, tailor_id, model_id, price, created_at, updated_at
		FROM tailor_models
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.TailorIDs; len(p) != 0 {
		whereQ += ` AND tailor_id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}
	if p := param.ModelIDs; len(p) != 0 {
		whereQ += ` AND model_id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", selectQ, whereQ), 1)

	rows, err := model.Toolkit.DB.Query(q, whereP...)
	if err != nil {
		return *new([]entity.TailorModelEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.TailorModelEntity)

		if err := rows.Scan(&t.ID, &t.TailorID, &t.ModelID, &t.Price, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.TailorModelEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
}

func (model *TailorModel) InsertBulk(param []InsertBulkTailorModelParam) (int, error) {
	if len(param) <= 0 {
		return 0, nil
	}

	trx, err := model.Toolkit.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer trx.Rollback()

	q := `INSERT INTO tailor_models (tailor_id, model_id, price) VALUES`
	p := []interface{}{}
	for i, v := range param {
		q += ` (?, ?, ?)`
		p = append(p, v.TailorID, v.ModelID, v.Price)

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
