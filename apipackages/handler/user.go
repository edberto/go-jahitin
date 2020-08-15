package handler

import (
	"fmt"
	"go-jahitin/apipackages"
	"go-jahitin/apipackages/request"
	"go-jahitin/apipackages/usecase"
	"log"
	"net/http"

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
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Invalid body"))
		log.Print(errors.Wrap(err, "Failed to bind request body"))
		return
	}

	if err := h.Toolkit.Validator.Struct(req); err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Missing information: %v", err))
		log.Print(errors.Wrap(err, "Missing information"))
		return
	}

	res, err := h.UserUC.Register((usecase.RegisterUserParam)(*req))
	if err != nil {
		switch err {
		default:
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Internal Server Error"))
			log.Print(errors.Wrap(err, "Failed to register user"))
			return
		}
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *User) GetOne(c *gin.Context) {
	return
}
