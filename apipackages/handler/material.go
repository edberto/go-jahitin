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
	IMaterial interface {
		GetAll(c *gin.Context)
	}

	Material struct {
		Toolkit *apipackages.Toolkit
		MaterialUC usecase.IMaterial
	}
)

func NewMaterialHandler(tk *apipackages.Toolkit) IMaterial {
	return &Material{
		Toolkit: tk,
		MaterialUC: usecase.NewMaterialUC(tk),
	}
}

func (h *Material) GetAll(c *gin.Context) {
	res, err := h.MaterialUC.GetAll(usecase.GetAllMaterialParam{})
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