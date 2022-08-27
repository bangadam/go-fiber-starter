package seeds

import (
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(db *gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "articles",
			Run: func(db *gorm.DB) error {
				return ArticleSeed(db)
			},
		},
	}
}
