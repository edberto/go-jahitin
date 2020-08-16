package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/model"
)

type (
	ITailorModel interface {
		InsertBulk(param []InsertBulkTailorModelParam) (res int, err error)
	}

	TailorModel struct {
		Toolkit          *apipackages.Toolkit
		TailorModelModel model.ITailorModel
	}

	InsertBulkTailorModelParam struct {
		TailorID int
		ModelID  int
		Price    float64
	}
)

func NewTailorModelUC(tk *apipackages.Toolkit) ITailorModel {
	return &TailorModel{
		Toolkit:          tk,
		TailorModelModel: model.NewTailorModelModel(tk),
	}
}

func (uc *TailorModel) InsertBulk(param []InsertBulkTailorModelParam) (res int, err error) {
	insertBulkTailorModelParam := []model.InsertBulkTailorModelParam{}
	for _, v := range param {
		insertBulkTailorModelParam = append(insertBulkTailorModelParam, model.InsertBulkTailorModelParam{
			TailorID: v.TailorID,
			ModelID:  v.ModelID,
			Price:    v.Price,
		})
	}

	return uc.TailorModelModel.InsertBulk(insertBulkTailorModelParam)
}
