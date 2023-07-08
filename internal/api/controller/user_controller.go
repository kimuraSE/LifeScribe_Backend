package controller

import (
	"LifeScribe_Backend/internal/api/model"
	"LifeScribe_Backend/internal/api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Login(c echo.Context) error
	SignUp(c echo.Context) error
	Destory(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (u *userController) Login(c echo.Context) error {
	req := model.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := model.User{
		ID : req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: []byte(req.Password),
	}

	token, err := u.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure = true
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (u *userController) SignUp(c echo.Context) error {
	req := model.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := model.User{
		ID : req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: []byte(req.Password),
	}

	token, err := u.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure = true
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (u *userController) Destory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	if err := u.uu.Destory(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure=true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (u *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK,
		echo.Map{
			"csrf_token": token,
		})
}
