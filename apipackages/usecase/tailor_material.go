package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/model"
)

type (
	ITailorMaterial interface {
		InsertBulk(param []InsertBulkTailorMaterialParam) (res int, err error)
	}

	TailorMaterial struct {
		Toolkit *apipackages.Toolkit
		TailorMaterialModel model.ITailorMaterial
	}

	InsertBulkTailorMaterialParam struct {
		TailorID   int
		MaterialID int
		Price      float64
	}
)

func NewTailorMaterialUC(tk *apipackages.Toolkit) ITailorMaterial {
	return &TailorMaterial{
		Toolkit: tk,
		TailorMaterialModel: model.NewTailorMaterialModel(tk),
	}
}

func (uc *TailorMaterial) InsertBulk(param []InsertBulkTailorMaterialParam) (res int, err error) {
	insertBulkTailorMaterialParam := []model.InsertBulkTailorMaterialParam{}
	for _, v := range param {
		insertBulkTailorMaterialParam = append(insertBulkTailorMaterialParam, model.InsertBulkTailorMaterialParam{
			TailorID: v.TailorID,
			MaterialID: v.MaterialID,
			Price: v.Price,
		})
	}

	return uc.TailorMaterialModel.InsertBulk(insertBulkTailorMaterialParam)
}