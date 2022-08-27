package schema

type Article struct {
	ID      uint64 `gorm:"primary_key" json:"id"`
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Base
}
