package model

import (
	"time"

	"main.go/global"
)

type task struct {
	Id     int
	Name   string
	Dueday time.Time
	Level  global.ImportanceType
}

func NewTask(id int, name string, dueday time.Time, level global.ImportanceType) *task {
	return &task{
		Id:     id,
		Name:   name,
		Dueday: dueday,
		Level:  level,
	}
}
