package service

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"src/model"
	"src/repository"
	"strconv"
)

func fetchRecords(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read plan input file "+filePath, err)
	}
	defer f.Close()

	//skip first line
	row1, _ := bufio.NewReader(f).ReadSlice('\n')
	f.Seek(int64(len(row1)), io.SeekStart)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse plan input file as CSV for "+filePath, err)
	}
	return records
}

func IngestPlansFromCsvFile() {
	records := fetchRecords("C:\\Users\\anjan\\OneDrive\\Desktop\\Karun\\Data\\Plan_data.csv")
	plans := []model.Plan{}
	for _, record := range records {
		plan := model.Plan{}
		plan.Id = record[0]
		plan.Duration, _ = strconv.Atoi(record[1])
		temp, _ := strconv.ParseFloat(record[2], 32)
		plan.Cost = float32(temp)
		plan.Sms, _ = strconv.Atoi(record[3])
		plan.Talktime = record[4]
		plan.Internet = record[5]
		plans = append(plans, plan)
	}
	repository.AddPlan(plans)
}
