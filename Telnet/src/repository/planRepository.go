package repository

import (
	"fmt"
	"src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func DbConnect() {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection established with database")
	}
}

func DbClose() {
	Dbase, _ := Db.DB()
	err := Dbase.Close()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection closed successfully")
	}
}

func AddPlan(plans []model.Plan) {
	DbConnect()
	if len(plans) < 10 {
		Db.Create(&plans)
	} else {
		Db.CreateInBatches(&plans, 3)
	}
	DbClose()
}

func GetPlan(plan model.Plan) model.Plan {
	Db.Find(&plan)
	return plan
}

func UpdatePlan(plan model.Plan) {
	Db.Save(plan)
}

func UpdatePlanDuration(plan model.Plan, duration int) {
	Db.Model(&plan).Update("Duration", duration)
}

func UpdatePlanDurationUsingId(id string, duration int) {
	Db.Model(model.Plan{}).Where("Id = ?", id).Update("Duration", duration)
}

func DeletePlanUsingId(id string) {
	plan := model.Plan{Id: id}
	Db.Delete(&plan)
	//delete using primary key
	// Db.Delete(&model.Plan,id)
}

func DeleteConsumerUsingDuration(duration int) {
	Db.Where("Duration = ?", duration).Delete(&model.Plan{})
}
