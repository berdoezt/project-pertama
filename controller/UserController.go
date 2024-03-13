package controller

import (
	"net/http"
	"project-pertama/model"
	"project-pertama/repository"
	"project-pertama/util"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepository repository.IUserRepository
}

func NewUserController(userRepository repository.IUserRepository) *userController {
	return &userController{
		userRepository: userRepository,
	}
}

func (uc *userController) Login(ctx *gin.Context) {
	var requestedUser model.User
	err := ctx.ShouldBindJSON(&requestedUser)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	user, err := uc.userRepository.GetByUsername(requestedUser.Username)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, r)
		return
	}

	if !util.HashMatched([]byte(user.Password), []byte(requestedUser.Password)) {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	token, err := util.GenerateJWTToken(user.IsAdmin, user.UUID)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, gin.H{
		"token": token,
	}, ""))
}

func (uc *userController) Register(ctx *gin.Context) {
	var newUser model.User
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	hashedPassword, err := util.Hash([]byte(newUser.Password))
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	newUser.Password = string(hashedPassword)
	createdUser, err := uc.userRepository.Create(newUser)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, gin.H{
		"uuid": createdUser.UUID,
	}, ""))
}
