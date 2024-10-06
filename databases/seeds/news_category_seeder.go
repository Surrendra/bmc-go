package seeds

import (
	"BaliMediaCenter/models"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

func SeedNewsCategory(db *gorm.DB) {
	newsCategories := []models.NewsCategory{
		{
			Code:        "",
			Name:        "General Daily News",
			Description: "General Daily News",
			Image:       "/image/naruto.jpg",
		},
		{
			Code:        "",
			Name:        "Weather",
			Description: "Weather",
			Image:       "/image/hinata-hp.png",
		},
	}

	user := models.User{}
	db.Where("username = ?", "surrendra").First(&user)
	for _, newsCategory := range newsCategories {
		NewNewsCategory := newsCategory
		NewNewsCategory.Slug = strings.ReplaceAll(newsCategory.Name, "-", " ")
		NewNewsCategory.Code = uuid.NewString()
		NewNewsCategory.CreatedUserName = user.Name
		NewNewsCategory.CreatedUserId = user.ID
		fmt.Println(NewNewsCategory)

		db.FirstOrCreate(&NewNewsCategory, models.NewsCategory{Name: NewNewsCategory.Name})
	}

}
