package seeds

import (
	"BaliMediaCenter/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {

	password, _ := bcrypt.GenerateFromPassword([]byte("11235811"), bcrypt.DefaultCost)
	users := []models.User{
		{
			Code:     "123012310",
			Name:     "Surendra Made",
			Email:    "surrendra@sunseeker.com",
			Username: "surrendra",
			Password: string(password),
		},
		{
			Code:     "991238912831s",
			Name:     "John Doe",
			Email:    "john@sunseeker.com",
			Username: "john",
			Password: string(password),
		},
	}

	for _, user := range users {
		newUser := user
		db.FirstOrCreate(&newUser, models.User{Code: user.Code})
	}

}
