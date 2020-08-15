package handler

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type (
	ITailor interface {
		GetAll(c *gin.Context)
	}

	Tailor struct {
		Toolkit  *apipackages.Toolkit
		TailorUC usecase.ITailor
	}
)

func NewTailorHandler(tk *apipackages.Toolkit) ITailor {
	return &Tailor{
		Toolkit:  tk,
		TailorUC: usecase.NewTailorUC(tk),
	}
}

func (h *Tailor) GetAll(c *gin.Context) {
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

	idsA := c.QueryArray("id")
	ids := []int{}
	for _, i := range idsA {
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{"Invalid Tailor ID"})
			log.Print(errors.Wrap(err, "Failed to convert Tailor ID"))
			return
		}

		ids = append(ids, id)
	}

	materialIDsA := c.QueryArray("material_id")
	materialIDs := []int{}
	for _, i := range materialIDsA {
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{"Invalid Material ID"})
			log.Print(errors.Wrap(err, "Failed to convert Material ID"))
			return
		}

		materialIDs = append(materialIDs, id)
	}

	modelIDsA := c.QueryArray("model_id")
	modelIDs := []int{}
	for _, i := range modelIDsA {
		id, err := strconv.Atoi(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error{"Invalid Model ID"})
			log.Print(errors.Wrap(err, "Failed to Model ID"))
			return
		}

		modelIDs = append(modelIDs, id)
	}

	res, err := h.TailorUC.GetAll(usecase.GetAllTailorParam{
		UserIDs:     userIDs,
		IDs:         ids,
		MaterialIDs: materialIDs,
		ModelIDs:    modelIDs,
		Keyword:     c.Query("keyword"),
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
