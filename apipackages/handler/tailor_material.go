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
	ITailorMaterial interface {
		InsertBulk(c *gin.Context)
	}

	TailorMaterial struct {
		Toolkit          *apipackages.Toolkit
		TailorMaterialUC usecase.ITailorMaterial
	}
)

func NewTailorMaterialHandler(tk *apipackages.Toolkit) ITailorMaterial {
	return &TailorMaterial{
		Toolkit:          tk,
		TailorMaterialUC: usecase.NewTailorMaterialUC(tk),
	}
}

func (h *TailorMaterial) InsertBulk(c *gin.Context) {
	req := new([]request.InsertBulkTailorMaterial)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid body"})
		log.Print(errors.Wrap(err, "Invalid request body"))
		return
	}

	param := []usecase.InsertBulkTailorMaterialParam{}
	for _, v := range *req {
		if err := h.Toolkit.Validator.Struct(v); err != nil {
			c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Invalid information: %v", err)})
			log.Print(errors.Wrap(err, "Invalid information"))
			return
		}

		param = append(param, usecase.InsertBulkTailorMaterialParam{
			TailorID:   v.TailorID,
			MaterialID: v.MaterialID,
			Price:      v.Price,
		})
	}

	res, err := h.TailorMaterialUC.InsertBulk(param)
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
