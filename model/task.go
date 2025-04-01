package model

import (
	"main.go/global"
)

type Task struct {
	Id       string // 生成uuid
	Name     string
	Deadline string
	Level    global.ImportanceType
}
