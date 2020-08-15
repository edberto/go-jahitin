package handler

import (
	"database/sql"
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/request"
	"go-jahitin/apipackages/usecase"
	"go-jahitin/helper/constants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type (
	ISession interface {
		Login(c *gin.Context)
		Logout(c *gin.Context)
		Refresh(c *gin.Context)
	}

	Session struct {
		SessionUC usecase.ISession
		Toolkit   *apipackages.Toolkit
	}
)

func NewSessionHandler(tk *apipackages.Toolkit) ISession {
	return &Session{
		SessionUC: usecase.NewSessionUC(tk),
		Toolkit:   tk,
	}
}

func (h *Session) Login(c *gin.Context) {
	req := new(request.LoginRequest)

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid body"})
		log.Print(errors.Wrap(err, "Failed to bind request body"))
		return
	}

	if err := h.Toolkit.Validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, Error{fmt.Sprintf("Invalid information: %v", err)})
		log.Print(errors.Wrap(err, "Missing information"))
		return
	}

	res, err := h.SessionUC.Login((usecase.LoginSessionParam)(*req))
	if err != nil {
		switch err {
		case sql.ErrNoRows, constants.ErrIncorrectPassword:
			c.JSON(http.StatusBadRequest, Error{"Username/Password incorrect!"})
			log.Print(errors.Wrap(err, "Failed login"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to login"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *Session) Logout(c *gin.Context) {
	userID := c.Request.Context().Value("userID").(int)

	err := h.SessionUC.Logout(usecase.LogoutSessionParam{
		UserID: userID,
	})
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to logout"))
			return
		}
	}

	c.JSON(http.StatusOK, nil)
	return
}

func (h *Session) Refresh(c *gin.Context) {
	claims, err := h.Toolkit.RefreshAuth.ExtractClaims(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid Token!"})
		log.Print(errors.Wrap(err, "Failed to extract claims"))
		return
	}

	res, err := h.SessionUC.Refresh(usecase.RefreshSessionParam{
		RefreshUUID: claims["refresh_uuid"].(string),
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusBadRequest, Error{"Invalid Token!"})
			log.Print(errors.Wrap(err, "Token does not exist"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to refresh token"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}
