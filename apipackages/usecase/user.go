package usecase

import (
	"database/sql"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/entity"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
	"go-jahitin/helper/constants"
	"go-jahitin/helper/uuid"

	"golang.org/x/crypto/bcrypt"
)

type (
	IUser interface {
		Register(param RegisterUserParam) (viewmodel.UserVM, error)
		GetOne(param GetOneUserParam) (viewmodel.UserVM, error)
	}

	User struct {
		UserModel   model.IUser
		TailorModel model.ITailor
		Toolkit     *apipackages.Toolkit
	}

	RegisterUserParam struct {
		Address  string
		Email    string
		Phone    string
		Name     string
		Username string
		Password string
		Role     string
	}
	GetOneUserParam struct {
		ID int
	}
)

func NewUserUC(tk *apipackages.Toolkit) IUser {
	return &User{
		UserModel:   model.NewUserModel(tk),
		TailorModel: model.NewTailorModel(tk),
		Toolkit:     tk,
	}
}

func (uc *User) Register(param RegisterUserParam) (viewmodel.UserVM, error) {
	user, err := uc.UserModel.GetOne(model.GetOneUserParam{
		Username: param.Username,
	})
	if err != nil && err != sql.ErrNoRows {
		return *new(viewmodel.UserVM), err
	}
	if user.ID != 0 {
		return *new(viewmodel.UserVM), constants.ErrUsernameHasBeenUsed
	}

	pwd := []byte(param.Password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return *new(viewmodel.UserVM), err
	}

	tailor, err := entity.TailorEntity{}, nil
	role := constants.RoleAtoI[param.Role]
	if role == constants.TailorRoleI {
		tailor, err = uc.TailorModel.InsertOne(model.InsertOneTailorParam{
			Address: param.Address,
			Email:   param.Email,
			Name:    param.Name,
			UUID:    uuid.NewUUID(),
		})
		if err != nil {
			return *new(viewmodel.UserVM), err
		}
	}

	user, err = uc.UserModel.InsertOne(model.InsertOneUserParam{
		TailorID: tailor.ID,
		Address:  param.Address,
		Email:    param.Email,
		Name:     param.Name,
		Password: string(hash),
		Role:     role,
		Username: param.Username,
		UUID:     uuid.NewUUID(),
	})
	if err != nil {
		return *new(viewmodel.UserVM), err
	}

	return viewmodel.UserVM{
		ID:       user.ID,
		TailorID: int(user.TailorID.Int32),
		Address:  user.Address,
		Email:    user.Email,
		Name:     user.Name,
		Password: param.Password,
		Role:     constants.RoleItoA[user.Role],
		Username: user.Username,
		UUID:     user.UUID,
	}, nil
}

func (uc *User) GetOne(param GetOneUserParam) (viewmodel.UserVM, error) {
	res := new(viewmodel.UserVM)

	user, err := uc.UserModel.GetOne(model.GetOneUserParam{
		ID: param.ID,
	})
	if err != nil {
		return *res, err
	}

	return viewmodel.UserVM{
		ID:       user.ID,
		TailorID: int(user.TailorID.Int32),
		Address:  user.Address,
		Email:    user.Email,
		Name:     user.Name,
		Role:     constants.RoleItoA[user.Role],
		Username: user.Username,
		UUID:     user.UUID,
	}, nil
}
