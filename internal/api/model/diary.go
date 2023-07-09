package model

import "time"

type DiaryRequest struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	EntryDate string `json:"entry_date"`
	Content   string `json:"content"`
}

type Diary struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	EntryDate time.Time `json:"entry_date"`
	Content   string    `json:"content"`
}
