package usecase

import (
	"LifeScribe_Backend/internal/api/model"
	"LifeScribe_Backend/internal/api/repository"
)

type IDiaryUsecase interface {
	CreateDiary(diary model.Diary) (model.Diary, error)
	GetDiary(diary model.Diary) (model.Diary, error)
	GetDiaries(userId uint) ([]model.Diary, error)
	UpdateDiary(diary model.Diary, userId uint, diaryId uint) (model.Diary, error)
	DeleteDiary(userId uint, diaryId uint) error
}

type diaryUsecase struct {
	dr repository.IDiaryRepository
}

func NewDiaryUsecase(dr repository.IDiaryRepository) IDiaryUsecase {
	return &diaryUsecase{dr}
}

func (d *diaryUsecase) CreateDiary(diary model.Diary) (model.Diary, error) {

	if err := d.dr.CreateDiary(&diary); err != nil {
		return model.Diary{}, err
	}

	res := model.Diary{
		ID:        diary.ID,
		UserID:    diary.UserID,
		EntryDate: diary.EntryDate,
		Content:   diary.Content,
	}

	return res, nil
}

func (d *diaryUsecase) GetDiary(diary model.Diary) (model.Diary, error) {
	if err := d.dr.GetDiary(&diary); err != nil {
		return model.Diary{}, err
	}

	res := model.Diary{
		ID:        diary.ID,
		UserID:    diary.UserID,
		EntryDate: diary.EntryDate,
		Content:   diary.Content,
	}

	return res, nil
}

func (d *diaryUsecase) GetDiaries(userId uint) ([]model.Diary, error) {

	diaries := []model.Diary{}
	if err := d.dr.GetDiaries(&diaries, userId); err != nil {
		return []model.Diary{}, err
	}

	res := []model.Diary{}

	for _, diary := range diaries {
		v := model.Diary{
			ID:        diary.ID,
			UserID:    diary.UserID,
			EntryDate: diary.EntryDate,
			Content:   diary.Content,
		}
		res = append(res, v)
	}
	return res, nil
}

func (d *diaryUsecase) UpdateDiary(diary model.Diary, userId uint, diaryId uint) (model.Diary, error) {
	if err := d.dr.UpdateDiary(&diary, userId, diaryId); err != nil {
		return model.Diary{}, err
	}

	res := model.Diary{
		ID:        diary.ID,
		UserID:    diary.UserID,
		EntryDate: diary.EntryDate,
		Content:   diary.Content,
	}

	return res, nil
}

func (d *diaryUsecase) DeleteDiary(userId uint, diaryId uint) error {
	if err := d.dr.DeleteDiary(userId, diaryId); err != nil {
		return err
	}
	return nil
}
