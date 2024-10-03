package seeds

import (
	"BaliMediaCenter/models"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	users := []models.User{
		{
			Code:     "123012310",
			Name:     "Surendra Made",
			Email:    "surrendra@sunseeker.com",
			Username: "surrendra",
			Password: "11235811",
		},
		{
			Code:     "991238912831s",
			Name:     "John Doe",
			Email:    "john@sunseeker.com",
			Username: "john",
			Password: "11235811",
		},
	}

	for _, user := range users {
		newUser := user
		db.FirstOrCreate(&newUser, models.User{Code: user.Code})
	}

}
