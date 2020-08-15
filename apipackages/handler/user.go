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
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pkg/errors"
)

type (
	IUser interface {
		Register(c *gin.Context)
		GetOne(c *gin.Context)
	}

	User struct {
		UserUC  usecase.IUser
		Toolkit *apipackages.Toolkit
	}
)

func NewUserHandler(tk *apipackages.Toolkit) IUser {
	return &User{
		UserUC:  usecase.NewUserUC(tk),
		Toolkit: tk,
	}
}

func (h *User) Register(c *gin.Context) {
	req := new(request.Register)

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

	res, err := h.UserUC.Register((usecase.RegisterUserParam)(*req))
	if err != nil {
		switch err {
		case constants.ErrUsernameHasBeenUsed:
			c.JSON(http.StatusBadRequest, Error{"Username has been used"})
			log.Print(errors.Wrap(err, "User existed"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *User) GetOne(c *gin.Context) {
	tokenID := c.Request.Context().Value("userID").(int)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{"Invalid ID"})
		log.Print(errors.Wrap(err, "Failed to convert ID"))
		return
	}

	if tokenID != id {
		c.JSON(http.StatusUnauthorized, Error{"Invalid Token"})
		log.Print("Token ID and param ID does not match")
	}

	res, err := h.UserUC.GetOne(usecase.GetOneUserParam{
		ID: id,
	})
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusBadRequest, Error{"User not found!"})
			log.Print(errors.Wrap(err, "User not found"))
			return
		default:
			c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}
