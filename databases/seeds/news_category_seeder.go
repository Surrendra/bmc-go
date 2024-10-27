package seeds

import (
	"BaliMediaCenter/helpers"
	"BaliMediaCenter/models"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

func SeedNewsCategory(db *gorm.DB) {
	StorageHelper := helpers.NewStorageHelper()
	defaultImage := StorageHelper.GetNewsCategoryPath("default-image.png")
	newsCategories := []models.NewsCategory{
		{
			Code:        "",
			Name:        "General Daily News",
			Description: "General Daily News",
			Image:       defaultImage,
		},
		{
			Code:        "",
			Name:        "Weather",
			Description: "Weather",
			Image:       defaultImage,
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
		NewNewsCategory.CreatedAt = time.Now()
		NewNewsCategory.UpdatedAt = time.Now()
		fmt.Println(NewNewsCategory)

		db.FirstOrCreate(&NewNewsCategory, models.NewsCategory{Name: NewNewsCategory.Name})
	}

}
