package repository

import (
	"src/model"

	"gorm.io/gorm"
)

func AddConsumer(consumers []model.Consumer) {
	dbConnect()
	if len(consumers) < 100 {
		Db.Create(&consumers)
	} else {
		Db.CreateInBatches(&consumers, 50)
	}
	dbClose()
}

func GetConsumer(db *gorm.DB, consumer model.Consumer) model.Consumer {
	db.Find(&consumer)
	return consumer
}

func UpdateConsumer(db *gorm.DB, consumer model.Consumer) {
	db.Save(consumer)
}

func UpdatePhoneNumber(db *gorm.DB, consumer model.Consumer, phoneNumber string) {
	db.Model(&consumer).Update("Phone_number", phoneNumber)
}

func findConsumerById(db *gorm.DB, consumer model.Consumer) model.Consumer {
	db.Find(&consumer, consumer.Id)
	return consumer
}

func findConsumerByName(db *gorm.DB, consumerNames []string) []model.Consumer {
	consumers := []model.Consumer{}
	db.Where("name in ?", consumerNames).Find(&consumers)
	return consumers
}

func UpdateConsumerDuration(db *gorm.DB, id string, duration int) {
	db.Model(model.Plan{}).Where("Id = ?", id).Update("Duration", duration)
}

func DeleteConsumerUsingId(db *gorm.DB, id int) {
	consumer := model.Consumer{Id: id}
	db.Delete(&consumer)
}
