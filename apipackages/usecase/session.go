package usecase

import (
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/model"
	"go-jahitin/apipackages/viewmodel"
	"go-jahitin/helper/constants"
	"go-jahitin/helper/uuid"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type (
	ISession interface {
		Login(param LoginSessionParam) (viewmodel.LoginVM, error)
		Logout(param LogoutSessionParam) error
		Refresh(param RefreshSessionParam) (viewmodel.RefreshVM, error)
	}

	Session struct {
		SessionModel   model.ISession
		UserModel      model.IUser
		UserTokenModel model.IUserToken
		Toolkit        *apipackages.Toolkit
	}

	LoginSessionParam struct {
		Username string
		Password string
	}
	LogoutSessionParam struct {
		UserID int
	}
	RefreshSessionParam struct {
		RefreshUUID string
	}
)

func NewSessionUC(tk *apipackages.Toolkit) ISession {
	return &Session{
		SessionModel:   model.NewSessionModel(tk),
		UserModel:      model.NewUserModel(tk),
		UserTokenModel: model.NewUserTokenModel(tk),
		Toolkit:        tk,
	}
}

func (uc *Session) Login(param LoginSessionParam) (viewmodel.LoginVM, error) {
	user, err := uc.UserModel.GetOne(model.GetOneUserParam{
		Username: param.Username,
	})
	if err != nil {
		return *new(viewmodel.LoginVM), err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(param.Password)); err != nil {
		return *new(viewmodel.LoginVM), constants.ErrIncorrectPassword
	}

	now := time.Now()

	accessAuth, accessUUID, accessExpiredAt := uc.Toolkit.AccessAuth, uuid.NewUUID(), now.Add(365*24*60*60*time.Second)
	_, err = uc.UserTokenModel.InsertOne(model.InsertOneUserTokenParam{
		UserID:    user.ID,
		UUID:      accessUUID,
		ExpiredAt: accessExpiredAt,
	})
	if err != nil {
		return *new(viewmodel.LoginVM), err
	}

	accessAuth.SetClaim("user_id", user.ID)
	accessAuth.SetClaim("access_uuid", accessUUID)
	accessAuth.SetClaim("expired_at", accessExpiredAt)
	accessAuth.SetClaim("role", constants.RoleItoA[user.Role])
	accessToken, err := accessAuth.CreateToken()
	if err != nil {
		return *new(viewmodel.LoginVM), err
	}

	refreshAuth, refreshUUID, refreshExpiredAt := uc.Toolkit.RefreshAuth, uuid.NewUUID(), now.Add(365*24*60*60*time.Second)
	_, err = uc.UserTokenModel.InsertOne(model.InsertOneUserTokenParam{
		UserID:    user.ID,
		UUID:      refreshUUID,
		ExpiredAt: refreshExpiredAt,
	})
	if err != nil {
		return *new(viewmodel.LoginVM), err
	}

	refreshAuth.SetClaim("user_id", user.ID)
	refreshAuth.SetClaim("refresh_uuid", refreshUUID)
	refreshAuth.SetClaim("expired_at", refreshExpiredAt)
	refreshToken, err := refreshAuth.CreateToken()
	if err != nil {
		return *new(viewmodel.LoginVM), err
	}

	return viewmodel.LoginVM{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *Session) Logout(param LogoutSessionParam) error {
	err := uc.UserTokenModel.DeleteAll(model.DeleteAllUserTokenParam{
		UserID: param.UserID,
	})

	return err
}

func (uc *Session) Refresh(param RefreshSessionParam) (viewmodel.RefreshVM, error) {
	userToken, err := uc.UserTokenModel.GetOne(model.GetOneUserTokenParam{
		UUID: param.RefreshUUID,
	})
	if err != nil {
		return *new(viewmodel.RefreshVM), err
	}

	user, err := uc.UserModel.GetOne(model.GetOneUserParam{
		ID: userToken.UserID,
	})
	if err != nil {
		return *new(viewmodel.RefreshVM), err
	}

	now := time.Now()

	accessAuth, accessUUID, accessExpiredAt := uc.Toolkit.AccessAuth, uuid.NewUUID(), now.Add(365*24*60*60*time.Second)
	_, err = uc.UserTokenModel.InsertOne(model.InsertOneUserTokenParam{
		UserID:    user.ID,
		UUID:      accessUUID,
		ExpiredAt: accessExpiredAt,
	})
	if err != nil {
		return *new(viewmodel.RefreshVM), err
	}

	accessAuth.SetClaim("user_id", user.ID)
	accessAuth.SetClaim("access_uuid", accessUUID)
	accessAuth.SetClaim("expired_at", accessExpiredAt)
	accessAuth.SetClaim("role", constants.RoleItoA[user.Role])
	accessToken, err := accessAuth.CreateToken()
	if err != nil {
		return *new(viewmodel.RefreshVM), err
	}

	return viewmodel.RefreshVM{
		AccessToken: accessToken,
	}, nil
}
