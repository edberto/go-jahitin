package model

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/helper"

	"github.com/lib/pq"
)

type (
	IOrder interface {
		GetAll(param GetAllOrderParam) ([]entity.OrderEntity, error)
		InsertOne(param InsertOneOrderParam) (entity.OrderEntity, error)
		UpdateStatusOne(param UpdateStatusOneOrderParam) (entity.OrderEntity, error)
	}

	Order struct {
		Toolkit *apipackages.Toolkit
	}

	GetAllOrderParam struct {
		UserIDs   []int
		TailorIDs []int
	}

	InsertOneOrderParam struct {
		UserID        int
		TailorID      int
		Status        int
		Price         float64
		UUID          string
		Specification entity.Specification
	}

	UpdateStatusOneOrderParam struct {
		ID     int
		Status int
	}
)

func NewOrderModel(tk *apipackages.Toolkit) IOrder {
	return &Order{
		Toolkit: tk,
	}
}

func (model *Order) GetAll(param GetAllOrderParam) ([]entity.OrderEntity, error) {
	res := new([]entity.OrderEntity)

	selectQ := `
		SELECT id, uuid, user_id, tailor_id, status, price, specification, created_at, updated_at
		FROM orders
	`

	whereQ := `WHERE deleted_at IS NULL`
	whereP := []interface{}{}
	if p := param.UserIDs; len(p) != 0 {
		whereQ += ` AND user_id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}
	if p := param.TailorIDs; len(p) != 0 {
		whereQ += ` AND tailor_id = ANY(?)`
		whereP = append(whereP, pq.Array(p))
	}

	q := helper.ReplacePlaceholder(fmt.Sprintf("%s%s", selectQ, whereQ), 1)

	rows, err := model.Toolkit.DB.Query(q, whereP...)
	if err != nil {
		return *new([]entity.OrderEntity), err
	}

	defer rows.Close()
	for rows.Next() {
		t := new(entity.OrderEntity)

		if err := rows.Scan(&t.ID, &t.UUID, &t.UserID, &t.TailorID, &t.Status, &t.Price, &t.Specification, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return *new([]entity.OrderEntity), err
		}

		*res = append(*res, *t)
	}

	return *res, err
}

func (model *Order) InsertOne(param InsertOneOrderParam) (entity.OrderEntity, error) {
	res := new(entity.OrderEntity)

	q := `
		INSERT INTO orders (user_id, tailor_id, status, price, uuid, specification)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, uuid, user_id, tailor_id, status, price, specification, created_at, updated_at
	`
	p := []interface{}{param.UserID, param.TailorID, param.Status, param.Price, param.UUID, param.Specification}

	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.ID, &res.UUID, &res.UserID, &res.TailorID, &res.Status, &res.Price, &res.Specification, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}

func (model *Order) UpdateStatusOne(param UpdateStatusOneOrderParam) (entity.OrderEntity, error) {
	res := new(entity.OrderEntity)

	q := `
		UPDATE orders
		SET status = ?, updated_at = NOW()
		WHERE id = ?
		RETURNING id, uuid, user_id, tailor_id, status, price, specification, created_at, updated_at
	`
	p := []interface{}{param.Status, param.ID}

	q = helper.ReplacePlaceholder(q, 1)
	err := model.Toolkit.DB.QueryRow(q, p...).Scan(&res.ID, &res.UUID, &res.UserID, &res.TailorID, &res.Status, &res.Price, &res.Specification, &res.CreatedAt, &res.UpdatedAt)

	return *res, err
}
