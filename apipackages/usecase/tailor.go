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
		MaterialModel       model.IMaterial
		ModelModel          model.IModel
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

	tailorModelMap := map[int][]entity.TailorModelEntity{}
	if p := param.ModelIDs; len(p) != 0 {
		tailorModels, err := uc.TailorModelModel.GetAll(model.GetAllTailorModelParam{
			ModelIDs: p,
		})
		if err != nil || len(tailorModels) == 0 {
			return *new([]viewmodel.TailorVM), err
		}

		for _, t := range tailorModels {
			mdl := tailorModelMap[t.TailorID]
			mdl = append(mdl, t)
			tailorModelMap[t.TailorID] = mdl
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

	tailorMaterialMap := map[int][]entity.TailorMaterialEntity{}
	if p := param.MaterialIDs; len(p) != 0 {
		tailorMaterials, err := uc.TailorMaterialModel.GetAll(model.GetAllTailorMaterialParam{
			MaterialIDs: p,
		})
		if err != nil || len(tailorMaterials) == 0 {
			return *new([]viewmodel.TailorVM), err
		}

		for _, t := range tailorMaterials {
			material := tailorMaterialMap[t.TailorID]
			material = append(material, t)
			tailorMaterialMap[t.TailorID] = material
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
			var ok bool
			for _, mt := range tailorMaterialMap[id] {
				if ok {
					break
				}
				for _, md := range tailorModelMap[id] {
					if ok {
						break
					}

					if (mt.Price + md.Price) <= param.Price {
						ok = true
					}
				}
			}
			if !ok {
				ids[i] = ids[len(ids)-1]
				ids = ids[:len(ids)-1]
			}
		}
	}
	if len(ids) == 0 {
		return *new([]viewmodel.TailorVM), nil
	}

	materials, err := uc.MaterialModel.GetAll(model.GetAllMaterialParam{
		IDs: param.MaterialIDs,
	})
	if err != nil {
		return *new([]viewmodel.TailorVM), err
	}

	materialMap := map[int]entity.MaterialEntity{}
	for _, m := range materials {
		materialMap[m.ID] = m
	}

	models, err := uc.ModelModel.GetAll(model.GetAllModelParam{
		IDs: param.ModelIDs,
	})
	if err != nil {
		return *new([]viewmodel.TailorVM), err
	}

	modelMap := map[int]entity.ModelEntity{}
	for _, m := range models {
		modelMap[m.ID] = m
	}

	tailors, err := uc.TailorModel.GetAll(model.GetAllTailorParam{
		IDs:     ids,
		Keyword: param.Keyword,
	})
	if err != nil {
		return *new([]viewmodel.TailorVM), err
	}

	for _, t := range tailors {
		for _, md := range tailorModelMap[t.ID] {
			for _, mt := range tailorMaterialMap[t.ID] {
				temp := viewmodel.TailorVM{
					ID:            t.ID,
					MaterialID:    mt.MaterialID,
					ModelID:       md.ModelID,
					MaterialName:  materialMap[mt.MaterialID].Name,
					MaterialColor: materialMap[mt.MaterialID].Color,
					ModelName:     modelMap[md.ModelID].Name,
					UUID:          t.UUID,
					Name:          t.Name.String,
					Phone:         t.Phone.String,
					Email:         t.Email,
					Address:       t.Address.String,
				}

				if md.Price != 0 {
					temp.Price += md.Price
				}

				if mt.Price != 0 {
					temp.Price += mt.Price
				}

				*res = append(*res, temp)
			}
		}
	}

	return *res, nil
}
