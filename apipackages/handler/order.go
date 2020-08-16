package handler

import (
	"database/sql"
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/request"
	"go-jahitin/apipackages/usecase"
	"go-jahitin/helper/constants"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type (
	IOrder interface {
		GetAll(c *gin.Context)
		GetOne(c *gin.Context)
		InsertOne(c *gin.Context)
		UpdateStatusOne(c *gin.Context)
	}

	Order struct {
		Toolkit *apipackages.Toolkit
		OrderUC usecase.IOrder
	}
)

func NewOrderHandler(tk *apipackages.Toolkit) IOrder {
	return &Order{
		Toolkit: tk,
		OrderUC: usecase.NewOrderUC(tk),
	}
}

func (h *Order) GetAll(c *gin.Context) {
	userIDsA := c.QueryArray("user_id")
	userIDs := []int{}
	for _, i := range userIDsA {
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{"Invalid User ID"})
			log.Print(errors.Wrap(err, "Failed to convert user ID"))
			return
		}

		userIDs = append(userIDs, id)
	}

	tailorIDsA := c.QueryArray("tailor_id")
	tailorIDs := []int{}
	for _, i := range tailorIDsA {
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{"Invalid Tailor ID"})
			log.Print(errors.Wrap(err, "Failed to convert tailor ID"))
			return
		}

		tailorIDs = append(tailorIDs, id)
	}

	idsA := c.QueryArray("id")
	ids := []int{}
	for _, i := range idsA {
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{"Invalid ID"})
			log.Print(errors.Wrap(err, "Failed to convert ID"))
			return
		}

		ids = append(ids, id)
	}

	statusA := c.QueryArray("user_id")
	status := []int{}
	for _, i := range statusA {
		v, e := constants.OrderStatusAtoI[i]
		if !e {
			c.JSON(http.StatusBadRequest, Error{"Invalid Status"})
			log.Print(errors.Wrap(nil, "Failed to convert Status"))
			return
		}

		status = append(status, v)
	}

	res, err := h.OrderUC.GetAll(usecase.GetAllOrderParam{
		IDs:       ids,
		Status:    status,
		UserIDs:   userIDs,
		TailorIDs: tailorIDs,
	})
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *Order) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(":id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid ID"})
		log.Print(errors.Wrap(err, "Invalid Order ID"))
		return
	}

	res, err := h.OrderUC.GetOne(usecase.GetOneOrderParam{
		ID: id,
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusBadRequest, Error{"Order does not exist"})
			log.Print(errors.Wrap(err, "Order does not exist"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *Order) InsertOne(c *gin.Context) {
	req := new(request.InsertOrder)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid body"})
		log.Print(errors.Wrap(err, "Invalid request body"))
		return
	}

	if err := h.Toolkit.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Invalid information: %v", err)})
		log.Print(errors.Wrap(err, "Invalid information"))
		return
	}

	res, err := h.OrderUC.InsertOne(usecase.InsertOneOrderParam{
		UserID: req.UserID,
		TailorID: req.TailorID,
		ModelID: req.ModelID,
		MaterialID: req.MaterialID,
		XSQty: req.XSQty,
		SQty: req.SQty,
		MQty: req.MQty,
		LQty: req.LQty,
		XLQty: req.XLQty,
		XXLQty: req.XXLQty,
		LLLQty: req.LLLQty,
		Price: req.Price,
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusBadRequest, Error{"Order does not exist"})
			log.Print(errors.Wrap(err, "Order does not exist"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}
	c.JSON(http.StatusOK, res)
	return
}

func (h *Order) UpdateStatusOne(c *gin.Context) {
	req := new(request.UpdateStatusOrder)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid body"})
		log.Print(errors.Wrap(err, "Invalid request body"))
		return
	}

	if err := h.Toolkit.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Invalid information: %v", err)})
		log.Print(errors.Wrap(err, "Invalid information"))
		return
	}

	res, err := h.OrderUC.UpdateStatusOne(usecase.UpdateStatusOneOrderParam{
		ID: req.ID,
		Status: constants.OrderStatusAtoI[req.Status],
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusBadRequest, Error{"Order does not exist"})
			log.Print(errors.Wrap(err, "Order does not exist"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}
