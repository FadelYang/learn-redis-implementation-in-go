package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/auth/dto"
	"project-root/modules/auth/repository"
	"project-root/modules/auth/service"
	userDTO "project-root/modules/users/dto"
	userService "project-root/modules/users/service"
	"project-root/tools"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
	userService userService.UserService
}

func NewAuthController(authService service.AuthService, userService userService.UserService) *AuthController {
	return &AuthController{
		authService: authService,
		userService: userService,
	}
}

// @Tags 					auth
// @Summary				Login User
// @Description 	Login user
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.LoginResponse]
// @Router				/auth/login [post]
// @Param					request body dto.LoginDTO true "request body for log in [RAW]"
func (c *AuthController) Login(ctx *gin.Context) {
	var login dto.LoginDTO
	if err := ctx.ShouldBindBodyWithJSON(&login); err != nil {
		log.Printf("failed to logged in: %v", err)

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": fmt.Sprintf("%s", err.Error()),
			},
		)
		return
	}

	accessToken, err := c.authService.Login(ctx, login)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": fmt.Sprintf("%s", err.Error()),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.LoginResponse]{
			Status:  http.StatusOK,
			Message: "log in success",
			Data:    accessToken,
		},
	)
}

// @Tags 					auth
// @Summary				Register User
// @Description 	Register user
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[userDTO.UserDTO]
// @Router				/auth/register [post]
// @Param					request body userDTO.CreateUser true "request body for register an user [RAW]"
func (c *AuthController) Register(ctx *gin.Context) {
	var user userDTO.CreateUser
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Printf("Failed to create user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", repository.ErrRegister, err.Error())})
		return
	}

	createdExample, err := c.userService.Create(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)

		var vErr *tools.ValidationError
		if errors.As(err, &vErr) {
			ctx.JSON(http.StatusConflict, vErr)
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", repository.ErrRegister, err.Error())})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		common.BaseResponse[userDTO.UserDTO]{
			Status:  http.StatusCreated,
			Message: "register successfully",
			Data:    createdExample,
		},
	)
}

// @Tags 					auth
// @Summary				Refresh Access Token
// @Description 	Refresh access token
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.RefreshResponse]
// @Router				/auth/refresh [post]
// @Param					request body dto.RefreshDTO true "request body for refresh access token [RAW]"
func (c *AuthController) Refresh(ctx *gin.Context) {
	var refresh dto.RefreshDTO
	if err := ctx.ShouldBindBodyWithJSON(&refresh); err != nil {
		log.Printf("failed to refresh access token: %v", err)

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": fmt.Sprintf("%s", err.Error()),
			},
		)
		return
	}

	accessToken, err := c.authService.RefreshLogin(ctx, refresh.RefreshToken)
	if err != nil {
		log.Printf("failed to refresh access token: %v", err)

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": fmt.Sprintf("%s", err.Error()),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.RefreshResponse]{
			Status:  http.StatusOK,
			Message: "success get refresh token",
			Data: dto.RefreshResponse{
				AccessToken: accessToken,
			},
		},
	)
}
