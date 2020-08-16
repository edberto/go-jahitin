package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
	"go-jahitin/helper/constants"
	"go-jahitin/helper/uuid"
)

type (
	IOrder interface {
		FindDetail(order entity.OrderEntity) (viewmodel.OrderVM, error)
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
		UserID     int
		TailorID   int
		ModelID    int
		MaterialID int
		XSQty      int
		SQty       int
		MQty       int
		LQty       int
		XLQty      int
		XXLQty     int
		LLLQty     int
		Price      float64
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
		user := userMap[o.UserID]
		tailor := tailorMap[o.TailorID]
		mdl := modelMap[o.Specification.ModelID]
		material := materialMap[o.Specification.MaterialID]

		t := viewmodel.OrderVM{
			ID:          o.ID,
			UserID:      user.ID,
			TailorID:    tailor.ID,
			UserName:    user.Name,
			UserPhone:   user.Phone,
			UserAddress: o.UserAddress.String,
			TailorName:  tailor.Name.String,
			TailorPhone: tailor.Phone.String,
			Status:      constants.OrderStatusItoA[o.Status],
			UUID:        o.UUID,
			Specification: viewmodel.Specification{
				MaterialID:     material.ID,
				ModelID:        mdl.ID,
				MaterialName:   material.Name,
				MaterialColor:  material.Color,
				MaterialDetail: material.Detail,
				ModelName:      mdl.Name,
				ModelDetail:    mdl.Detail,
				Quantity:       (viewmodel.Quantity)(o.Specification.Qty),
			},
		}

		*res = append(*res, t)
	}

	return *res, err
}

func (uc *Order) GetOne(param GetOneOrderParam) (viewmodel.OrderVM, error) {
	order, err := uc.OrderModel.GetOne(model.GetOneOrderParam{
		ID: param.ID,
	})
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	return uc.FindDetail(order)
}

func (uc *Order) InsertOne(param InsertOneOrderParam) (viewmodel.OrderVM, error) {
	order, err := uc.OrderModel.InsertOne(model.InsertOneOrderParam{
		UserID:        param.UserID,
		TailorID:      param.TailorID,
		Status:        constants.OrderStatusWaitingI,
		Price:         param.Price,
		UUID:          uuid.NewUUID(),
		Specification: entity.Specification {
			MaterialID: param.MaterialID,
			ModelID: param.ModelID,
			Qty: entity.Qty{
				XS: param.XSQty,
				S: param.SQty,
				M: param.MQty,
				L: param.LQty,
				XL: param.XLQty,
				XXL: param.XXLQty,
				LLL: param.LLLQty,
			},
		},
	})
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	return uc.FindDetail(order)
}

func (uc *Order) UpdateStatusOne(param UpdateStatusOneOrderParam) (viewmodel.OrderVM, error) {
	order, err := uc.OrderModel.UpdateStatusOne((model.UpdateStatusOneOrderParam)(param))
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	return uc.FindDetail(order)
}

func (uc *Order) FindDetail(order entity.OrderEntity) (viewmodel.OrderVM, error) {
	user, err := uc.UserModel.GetOne(model.GetOneUserParam{
		ID: order.UserID,
	})
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	tailor, err := uc.TailorModel.GetOne(model.GetOneTailorParam{
		ID: order.TailorID,
	})
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	mdl, err := uc.ModelModel.GetOne(model.GetOneModelParam{
		ID: order.Specification.ModelID,
	})
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	material, err := uc.MaterialModel.GetOne(model.GetOneMaterialParam{
		ID: order.Specification.MaterialID,
	})
	if err != nil {
		return *new(viewmodel.OrderVM), err
	}

	return viewmodel.OrderVM{
		ID:          order.ID,
		UserID:      user.ID,
		TailorID:    tailor.ID,
		UserName:    user.Name,
		UserPhone:   user.Phone,
		UserAddress: order.UserAddress.String,
		TailorName:  tailor.Name.String,
		TailorPhone: tailor.Phone.String,
		Status:      constants.OrderStatusItoA[order.Status],
		UUID:        order.UUID,
		Specification: viewmodel.Specification{
			MaterialID:     material.ID,
			ModelID:        mdl.ID,
			MaterialName:   material.Name,
			MaterialColor:  material.Color,
			MaterialDetail: material.Detail,
			ModelName:      mdl.Name,
			ModelDetail:    mdl.Detail,
			Quantity:       (viewmodel.Quantity)(order.Specification.Qty),
		},
	}, nil
}
