package seeds

import (
	"fmt"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	SeedUser(db)
	SeedNewsCategory(db)
	fmt.Println("Finish Running All Seed")
}
