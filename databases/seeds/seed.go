package seeds

import (
	"fmt"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	SeedUser(db)
	fmt.Println("Finish Running All Seed")
}
