package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
)

type (
	ITailor interface {
		GetAll(param GetAllTailorParam) ([]viewmodel.TailorVM, error)
	}

	Tailor struct {
		Toolkit             *apipackages.Toolkit
		TailorModel         model.ITailor
		TailorModelModel    model.ITailorModel
		TailorMaterialModel model.ITailorMaterial
		UserModel           model.IUser
	}

	GetAllTailorParam struct {
		UserIDs     []int
		IDs         []int
		MaterialIDs []int
		ModelIDs    []int
		Keyword     string
		Price       float64
	}
)

func NewTailorUC(tk *apipackages.Toolkit) ITailor {
	return &Tailor{
		Toolkit:             tk,
		TailorModel:         model.NewTailorModel(tk),
		TailorModelModel:    model.NewTailorModelModel(tk),
		TailorMaterialModel: model.NewTailorMaterialModel(tk),
		UserModel:           model.NewUserModel(tk),
	}
}

func (uc *Tailor) GetAll(param GetAllTailorParam) ([]viewmodel.TailorVM, error) {
	res := new([]viewmodel.TailorVM)

	ids := param.IDs
	if p := param.UserIDs; len(p) != 0 {
		users, err := uc.UserModel.GetAll(model.GetAllUserParam{
			IDs: p,
		})
		if err != nil || len(users) == 0 {
			return *new([]viewmodel.TailorVM), err
		}

		for _, u := range users {
			if u.TailorID.Valid {
				ids = append(ids, int(u.TailorID.Int32))
			}
		}
	}
	if len(ids) == 0 {
		return *new([]viewmodel.TailorVM), nil
	}

	tailorModelMap := map[int]entity.TailorModelEntity{}
	if p := param.ModelIDs; len(p) != 0 {
		tailorModels, err := uc.TailorModelModel.GetAll(model.GetAllTailorModelParam{
			ModelIDs: p,
		})
		if err != nil || len(tailorModels) == 0 {
			return *new([]viewmodel.TailorVM), err
		}

		for _, t := range tailorModels {
			tailorModelMap[t.TailorID] = t
		}
	}
	for i, id := range ids {
		if _, e := tailorModelMap[id]; !e {
			ids[i] = ids[len(ids)-1]
			ids = ids[:len(ids)-1]
		}
	}
	if len(ids) == 0 {
		return *new([]viewmodel.TailorVM), nil
	}

	tailorMaterialMap := map[int]entity.TailorMaterialEntity{}
	if p := param.MaterialIDs; len(p) != 0 {
		tailorMaterials, err := uc.TailorMaterialModel.GetAll(model.GetAllTailorMaterialParam{
			MaterialIDs: p,
		})
		if err != nil || len(tailorMaterials) == 0 {
			return *new([]viewmodel.TailorVM), err
		}

		for _, t := range tailorMaterials {
			tailorMaterialMap[t.TailorID] = t
		}
	}
	for i, id := range ids {
		if _, e := tailorMaterialMap[id]; !e {
			ids[i] = ids[len(ids)-1]
			ids = ids[:len(ids)-1]
		}
	}
	if len(ids) == 0 {
		return *new([]viewmodel.TailorVM), nil
	}

	if param.Price != 0 {
		for i, id := range ids {
			material := tailorMaterialMap[id]
			mdl := tailorModelMap[id]

			if (material.Price + mdl.Price) > param.Price {
				ids[i] = ids[len(ids)-1]
				ids = ids[:len(ids)-1]
			}
		}
	}
	if len(ids) == 0 {
		return *new([]viewmodel.TailorVM), nil
	}

	tailors, err := uc.TailorModel.GetAll(model.GetAllTailorParam{
		IDs:     ids,
		Keyword: param.Keyword,
	})
	if err != nil {
		return *new([]viewmodel.TailorVM), err
	}

	for _, t := range tailors {
		temp := viewmodel.TailorVM{
			ID:      t.ID,
			UUID:    t.UUID,
			Name:    t.Name.String,
			Phone:   t.Phone.String,
			Email:   t.Email,
			Address: t.Address.String,
		}

		if v, e := tailorModelMap[t.ID]; e {
			temp.Price += v.Price
		}

		if v, e := tailorModelMap[t.ID]; e {
			temp.Price += v.Price
		}

		*res = append(*res, temp)
	}

	return *res, nil
}
