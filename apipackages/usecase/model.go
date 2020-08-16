package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
)

type (
	IModel interface {
		GetAll(param GetAllModelParam) ([]viewmodel.ModelVM, error)
	}

	Model struct {
		Toolkit *apipackages.Toolkit
		ModelModel model.IModel
	}

	GetAllModelParam struct {}
)

func NewModelUC(tk *apipackages.Toolkit) IModel {
	return &Model{
		Toolkit: tk,
		ModelModel: model.NewModelModel(tk),
	}
}

func (uc *Model) GetAll(param GetAllModelParam) ([]viewmodel.ModelVM, error) {
	res := new([]viewmodel.ModelVM)

	models, err := uc.ModelModel.GetAll(model.GetAllModelParam{})
	if err != nil {
		return *new([]viewmodel.ModelVM), err
	}

	for _, m := range models {
		t := viewmodel.ModelVM {
			ID: m.ID,
			Name: m.Name,
			Detail: m.Detail,
		}

		*res = append(*res, t)
	}

	return *res, nil
}