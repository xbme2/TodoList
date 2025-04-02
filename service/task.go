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

// func createGoogleCalendarEvent() {
// 	// 加载认证信息
// 	ctx := context.Background()
// 	clientSecret := "path_to_your_credentials.json"

// 	b, err := os.ReadFile(clientSecret)
// 	if err != nil {
// 		log.Fatalf("Unable to read client secret file: %v", err)
// 	}

// 	// 获取 Google API 认证
// 	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
// 	if err != nil {
// 		log.Fatalf("Unable to parse client secret file to config: %v", err)
// 	}

// 	client := getClient(config)

// 	// 创建 Calendar 服务客户端
// 	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Calendar client: %v", err)
// 	}
// }

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
