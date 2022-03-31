package domain

type User struct {
	ID       int    `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}
