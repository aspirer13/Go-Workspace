package repository

import (
	"src/model"

	"gorm.io/gorm"
)

func DropTable(db *gorm.DB) {
	db.Migrator().DropTable(&model.Tester{})
}

func CreateTable(db *gorm.DB) {
	db.Migrator().CreateTable(&model.Tester{})
}
