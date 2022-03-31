package domain

type Note struct {
	ID          string `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
