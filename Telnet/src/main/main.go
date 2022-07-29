package main

import (
	"src/service"
)

func main() {
	service.IngestPlansFromCsvFile()
	service.InjestConsumersFromCsvFile()
}
