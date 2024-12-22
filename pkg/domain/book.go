package domain

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"unique;not null" copier:"must"`
	Name      string    `json:"name" copier:"must"`
	Author    string    `json:"author" copier:"must"`
	ISBNCode  string    `json:"isbnCode" copier:"must"`
	Score     int       `json:"score" gorm:"default:0;" copier:"must"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;" copier:"must"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoCreateTime;" copier:"must"`
}
