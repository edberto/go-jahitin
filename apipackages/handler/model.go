package handler

import (
	"log"
	"net/http"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/usecase"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type (
	IModel interface {
		GetAll(c *gin.Context)
	}

	Model struct {
		Toolkit *apipackages.Toolkit
		ModelUC usecase.IModel
	}
)

func NewModelHandler(tk *apipackages.Toolkit) IModel {
	return &Model{
		Toolkit: tk,
		ModelUC: usecase.NewModelUC(tk),
	}
}

func (h *Model) GetAll(c *gin.Context) {
	res, err := h.ModelUC.GetAll(usecase.GetAllModelParam{})
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