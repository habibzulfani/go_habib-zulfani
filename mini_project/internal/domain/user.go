package domain

// User represents a user in the system
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"` // Added field to identify admin users
}
