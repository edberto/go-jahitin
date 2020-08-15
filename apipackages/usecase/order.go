package usecase

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
)

type (
	IOrder interface {
		GetAll(param GetAllOrderParam) ([]viewmodel.OrderVM, error)
		GetOne(param GetOneOrderParam) (viewmodel.OrderVM, error)
		InsertOne(param InsertOneOrderParam) (viewmodel.OrderVM, error)
		UpdateStatusOne(param UpdateStatusOneOrderParam) (viewmodel.OrderVM, error)
	}

	Order struct {
		Toolkit       *apipackages.Toolkit
		OrderModel    model.IOrder
		UserModel     model.IUser
		TailorModel   model.ITailor
		MaterialModel model.IMaterial
		ModelModel    model.IModel
	}

	GetAllOrderParam struct {
		IDs       []int
		Status    []int
		UserIDs   []int
		TailorIDs []int
	}

	GetOneOrderParam struct {
		ID int
	}

	InsertOneOrderParam struct {
		UserID        int
		TailorID      int
		Status        int
		Price         float64
		UUID          string
		Specification entity.Specification
	}

	UpdateStatusOneOrderParam struct {
		ID     int
		Status int
	}
)

func NewOrderUC(tk *apipackages.Toolkit) IOrder {
	return &Order{
		Toolkit:       tk,
		OrderModel:    model.NewOrderModel(tk),
		UserModel:     model.NewUserModel(tk),
		TailorModel:   model.NewTailorModel(tk),
		MaterialModel: model.NewMaterialModel(tk),
		ModelModel:    model.NewModelModel(tk),
	}
}

func (uc *Order) GetAll(param GetAllOrderParam) ([]viewmodel.OrderVM, error) {
	res := new([]viewmodel.OrderVM)

	orders, err := uc.OrderModel.GetAll(model.GetAllOrderParam{
		IDs:       param.IDs,
		UserIDs:   param.UserIDs,
		TailorIDs: param.TailorIDs,
		Status:    param.Status,
	})
	if err != nil || len(orders) == 0 {
		return *new([]viewmodel.OrderVM), err
	}

	userIDs := []int{}
	tailorIDs := []int{}
	materialIDs := []int{}
	modelIDs := []int{}
	for _, o := range orders {
		userIDs = append(userIDs, o.UserID)
		tailorIDs = append(tailorIDs, o.TailorID)
		materialIDs = append(materialIDs, o.Specification.MaterialID)
		modelIDs = append(modelIDs, o.Specification.ModelID)
	}

	users, err := uc.UserModel.GetAll(model.GetAllUserParam{
		IDs: userIDs,
	})
	if err != nil || len(users) == 0 {
		return *new([]viewmodel.OrderVM), err
	}

	userMap := map[int]entity.UserEntity{}
	for _, u := range users {
		userMap[u.ID] = u
	}

	tailors, err := uc.TailorModel.GetAll(model.GetAllTailorParam{
		IDs: tailorIDs,
	})
	if err != nil || len(tailors) == 0 {
		return *new([]viewmodel.OrderVM), err
	}

	tailorMap := map[int]entity.TailorEntity{}
	for _, t := range tailors {
		tailorMap[t.ID] = t
	}

	models, err := uc.ModelModel.GetAll(model.GetAllModelParam{
		IDs: modelIDs,
	})
	if err != nil || len(models) == 0 {
		return *new([]viewmodel.OrderVM), err
	}

	modelMap := map[int]entity.ModelEntity{}
	for _, m := range models {
		modelMap[m.ID] = m
	}

	materials, err := uc.MaterialModel.GetAll(model.GetAllMaterialParam{
		IDs: materialIDs,
	})
	if err != nil || len(materials) == 0 {
		return *new([]viewmodel.OrderVM), err
	}

	materialMap := map[int]entity.MaterialEntity{}
	for _, m := range materials {
		materialMap[m.ID] = m
	}

	for _, o := range orders {
		t := viewmodel.OrderVM{}
		fmt.Println(o, t)
	}

	return *res, err
}

func (uc *Order) GetOne(param GetOneOrderParam) (viewmodel.OrderVM, error) {
	res := new(viewmodel.OrderVM)
	return *res, nil
}

func (uc *Order) InsertOne(param InsertOneOrderParam) (viewmodel.OrderVM, error) {
	res := new(viewmodel.OrderVM)
	return *res, nil
}

func (uc *Order) UpdateStatusOne(param UpdateStatusOneOrderParam) (viewmodel.OrderVM, error) {
	res := new(viewmodel.OrderVM)
	return *res, nil
}
