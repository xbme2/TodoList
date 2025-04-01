package service

import (
	"main.go/global"
	"main.go/model"
)

// new task类
func NewTask(id string, name string, deadline string, level global.ImportanceType) *model.Task {
	return &model.Task{
		Id:       id,
		Name:     name,
		Deadline: deadline,
		Level:    level,
	}
}

func AddTask(tasks []*model.Task, useCsv bool) {
	if useCsv {
		// 使用CSV存储数据
		for _, task := range tasks {
			//todo 并行优化
			WriteCSV(task, "test.csv")
		}
	} else {
		// 连接数据库存储
		// todo
	}
}

// func TasksToString(tasks []*model.Task) (stringTasks []string) {
// 	for _, t := range tasks {

// 	}
// }

// func TaskToString(task *model.Task) (stringTask string) {
// 	stringTask += string(task.Id)
// 	stringTask += task.Name

// 	stringTask += task.Dueday.Format(task.Dueday)
// 	stringTask += string(task.Level)

// }
