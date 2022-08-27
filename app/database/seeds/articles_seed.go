package seeds

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"gorm.io/gorm"
)

func ArticleSeed(db *gorm.DB) error {
	var seeders [][]string

	seeders = append(seeders,
		[]string{
			"title 1",
			"lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		},
	)

	for _, seeder := range seeders {
		if err := db.Create(&schema.Article{
			Title:   seeder[0],
			Content: seeder[1],
		}).Error; err != nil {
			return err
		}
	}

	return nil
}
