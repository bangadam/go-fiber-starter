package seeds

import (
	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"gorm.io/gorm"
)

type ArticleSeeder struct{}

var articles = []schema.Article{
	{
		Title:   "Example 1",
		Content: "Lorem Ipsum",
	},
	{
		Title:   "Example 2",
		Content: "Dolor Sit Amet",
	},
}

func (ArticleSeeder) Seed(conn *gorm.DB) error {
	if err := conn.CreateInBatches(&articles, 100).Error; err != nil {
		return err
	}

	return nil
}

func (ArticleSeeder) Count(conn *gorm.DB) (int, error) {
	var count int64
	if err := conn.Model(&schema.Article{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
