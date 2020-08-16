package handler

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/request"
	"go-jahitin/apipackages/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type (
	ITailorModel interface {
		InsertBulk(c *gin.Context)
	}

	TailorModel struct {
		Toolkit          *apipackages.Toolkit
		TailorModelUC usecase.ITailorModel
	}
)

func NewTailorModelHandler(tk *apipackages.Toolkit) ITailorModel {
	return &TailorModel{
		Toolkit:          tk,
		TailorModelUC: usecase.NewTailorModelUC(tk),
	}
}

func (h *TailorModel) InsertBulk(c *gin.Context) {
	req := new([]request.InsertBulkTailorModel)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid body"})
		log.Print(errors.Wrap(err, "Invalid request body"))
		return
	}

	param := []usecase.InsertBulkTailorModelParam{}
	for _, v := range *req {
		if err := h.Toolkit.Validator.Struct(v); err != nil {
			c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Invalid information: %v", err)})
			log.Print(errors.Wrap(err, "Invalid information"))
			return
		}

		param = append(param, usecase.InsertBulkTailorModelParam{
			TailorID:   v.TailorID,
			ModelID: v.ModelID,
			Price:      v.Price,
		})
	}

	res, err := h.TailorModelUC.InsertBulk(param)
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to insert tailor material"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return

}
