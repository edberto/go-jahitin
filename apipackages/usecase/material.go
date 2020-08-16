package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
)

type (
	IMaterial interface {
		GetAll(param GetAllMaterialParam) ([]viewmodel.MaterialVM, error)
	}

	Material struct {
		Toolkit *apipackages.Toolkit
		MaterialModel model.IMaterial
	}

	GetAllMaterialParam struct {}
)

func NewMaterialUC(tk *apipackages.Toolkit) IMaterial {
	return &Material{
		Toolkit: tk,
		MaterialModel: model.NewMaterialModel(tk),
	}
}

func (uc *Material) GetAll(param GetAllMaterialParam) ([]viewmodel.MaterialVM, error) {
	res := new([]viewmodel.MaterialVM)

	materials, err := uc.MaterialModel.GetAll(model.GetAllMaterialParam{})
	if err != nil {
		return *new([]viewmodel.MaterialVM), err
	}

	for _, m := range materials {
		t := viewmodel.MaterialVM {
			ID: m.ID,
			Name: m.Name,
			Color: m.Color,
			Detail: m.Detail,
		}

		*res = append(*res, t)
	}

	return *res, nil
}

