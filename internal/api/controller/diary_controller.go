package controller

import (
	"LifeScribe_Backend/internal/api/model"
	"LifeScribe_Backend/internal/api/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IDiaryController interface {
	Create(c echo.Context) error
	Read(c echo.Context) error
	AllRead(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type diaryController struct {
	du usecase.IDiaryUsecase
}

func NewDiaryController(du usecase.IDiaryUsecase) IDiaryController {
	return &diaryController{du}
}

func (d *diaryController) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	req := model.DiaryRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	parsedDate, err := time.Parse("2006-01-02", req.EntryDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	diary := model.Diary{
		ID:        req.ID,
		UserID:    uint(userId),
		EntryDate: parsedDate,
		Content:   req.Content,
	}

	res, err := d.du.CreateDiary(diary)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (d *diaryController) Read(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("id")
	diaryId, _ := strconv.Atoi(id)

	diary := model.Diary{
		ID:     uint(diaryId),
		UserID: uint(userId),
	}

	res, err := d.du.GetDiary(diary)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (d *diaryController) AllRead(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	diaries, err := d.du.GetDiaries(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, diaries)
}

func (d *diaryController) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	req := model.DiaryRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	diaryId, _ := strconv.Atoi(id)

	parsedDate, err := time.Parse("2006-01-02", req.EntryDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	diary := model.Diary{
		EntryDate: parsedDate,
		Content:   req.Content,
	}

	res, err := d.du.UpdateDiary(diary, uint(userId), uint(diaryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}

func (d *diaryController) Delete(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("id")
	diaryId, _ := strconv.Atoi(id)

	err := d.du.DeleteDiary(uint(userId), uint(diaryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Diary Deleted")
}
