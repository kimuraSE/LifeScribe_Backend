package repository

import (
	"LifeScribe_Backend/internal/api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDiaryRepository interface {
	CreateDiary(diary *model.Diary) error
	GetDiary(diary *model.Diary) error
	GetDiaries(diaries *[]model.Diary, userId uint) error
	UpdateDiary(diary *model.Diary, userId uint, diaryId uint) error
	DeleteDiary(userId uint, diaryId uint) error
}

type diaryRepository struct {
	db *gorm.DB
}

func NewDiaryRepository(db *gorm.DB) IDiaryRepository {
	return &diaryRepository{db}
}

func (d *diaryRepository) CreateDiary(diary *model.Diary) error {
	if err := d.db.Table("diaries").Create(diary).Error; err != nil {
		return err
	}
	return nil
}

func (d *diaryRepository) GetDiary(diary *model.Diary) error {
	if err := d.db.Table("diaries").Where("user_id = ? AND id = ?", diary.UserID, diary.ID).First(diary).Error; err != nil {
		return err
	}
	return nil
}

func (d *diaryRepository) GetDiaries(diaries *[]model.Diary, userId uint) error {
	if err := d.db.Table("diaries").Where("user_id = ?", userId).Find(diaries).Error; err != nil {
		return err
	}
	return nil
}

func (d *diaryRepository) UpdateDiary(diary *model.Diary, userId uint, diaryId uint) error {
	if err := d.db.Model(diary).Clauses(clause.Returning{}).Where("user_id = ? AND id = ?", userId, diaryId).Updates(
		map[string]interface{}{
			"entry_date": diary.EntryDate,
			"content":    diary.Content,
		},
	).Error; err != nil {
		return err
	}
	return nil
}

func (d *diaryRepository) DeleteDiary(userId uint, diaryId uint) error {
	if err := d.db.Table("diaries").Where("user_id = ? AND id = ?", userId, diaryId).Delete(&model.Diary{}).Error; err != nil {
		return err
	}
	return nil
}
