package domain

import "time"

// Book represents a book in the library
type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Published time.Time `json:"published"`
	UserID    uint      `json:"user_id"` // Added field to associate books with users
}