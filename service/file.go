package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"main.go/model"
)

// 此文件实现对文件的读写

func WriteCSV(task *model.Task, csvFile string) {
	file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	line := []string{
		task.Id,
		task.Name,
		task.Deadline,
		strconv.Itoa(int(task.Level)),
	}
	writer.Write(line)
}
