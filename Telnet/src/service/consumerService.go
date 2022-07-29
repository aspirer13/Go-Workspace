package service

import (
	"src/model"
	"src/repository"
	"strconv"
)

func InjestConsumersFromCsvFile() {
	records := fetchRecords("C:\\Users\\anjan\\OneDrive\\Desktop\\Karun\\Data\\Consumer_data.csv")
	consumers := []model.Consumer{}
	for _, record := range records {
		consumer := model.Consumer{}
		consumer.Name = record[0]
		consumer.Id, _ = strconv.Atoi(record[1])
		consumer.Age, _ = strconv.Atoi(record[2])
		consumer.PhoneNumber = record[3]
		consumer.PlanId = record[4]
		consumers = append(consumers, consumer)
	}
	repository.AddConsumer(consumers)
}
