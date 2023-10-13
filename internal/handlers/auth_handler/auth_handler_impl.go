package auth_handler

import (
	"auction-be/dtos/user_dtos"
	"auction-be/internal/errors/resterrors"
	"auction-be/internal/services/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthImpl struct {
	authService auth_service.Auth
}

func NewAuthHandler(auth auth_service.Auth) Auth {
	return &AuthImpl{authService: auth}
}

func (a AuthImpl) Login(c *gin.Context) {
	var loginRequest user_dtos.UserLoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, resterrors.CreateNew(http.StatusBadRequest, "VALIDATION__Invalid_Signup_Body", "Invalid json body"))
		return
	}

	token, restErr := a.authService.Login(&loginRequest)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (a AuthImpl) Signup(c *gin.Context) {
	var signupRequest user_dtos.SignUpRequest

	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(http.StatusBadRequest, resterrors.CreateNew(http.StatusBadRequest, "VALIDATION__Invalid_Signup_Body", "Invalid json body"))
		return
	}

	userResponse, restErr := a.authService.SignUp(&signupRequest)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}
