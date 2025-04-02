package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"main.go/global"
	"main.go/model"
	"main.go/service"
)

var addCommand = &cobra.Command{
	Use:   "add [-s] [-d]",
	Short: global.Iswroking,
	Long:  "thc command is used to add one or more tasks into your list",
	Run: func(cmd *cobra.Command, args []string) {

		names, _ := cmd.Flags().GetStringArray("name")
		deadlines, _ := cmd.Flags().GetStringArray("deadline")
		levels, _ := cmd.Flags().GetStringArray("level")
		for _, name := range names {
			fmt.Println(name)
		}
		for _, deadline := range deadlines {
			fmt.Println("-d " + deadline)
		}
		fmt.Println(len(deadlines))
		if err := checkArgs(names, deadlines, levels); err != nil {
			log.Fatalf(err.Error())
		}
		tasks, err := generateTasksByArgs(names, deadlines, levels)
		if err != nil {
			log.Fatalf(err.Error())
		}
		if tasks == nil {
			log.Fatalf(" tasks is nil")
		}
		service.AddTask(tasks, true)
	},
}

func checkArgs(names, deadlines, levels []string) error {
	if len(names) == 0 {
		return errors.New("请输入任务名称")
	} else if len(deadlines) == 0 {
		return errors.New("请输入任务截止日期")
	} else if len(names) > 1 && len(deadlines) > 1 && len(names) != len(deadlines) {
		return errors.New("请分配与任务相配对的截止日期")
	} else if len(names) > 1 && len(levels) > 1 && len(names) != len(levels) {
		return errors.New("请分配与任务相配对的优先级")
	}

	return nil
}

func generateTasksByArgs(names, deadlines, levels []string) ([]*model.Task, error) {
	tasks := make([]*model.Task, len(names))
	deadLen, imLen := len(deadlines), len(levels)
	fmt.Println(deadLen)
	for i, name := range names {
		deadline := deadlines[i%deadLen]
		level := global.Normal
		if imLen != 0 {
			val, err := strconv.Atoi(levels[i%imLen])
			if err != nil {
				return nil, err
			}
			level = global.ImportanceType(val)
		}
		task := service.NewTask(uuid.NewString(), name, deadline, level)
		if task == nil { // 避免 nil
			return nil, fmt.Errorf("任务创建失败: %s", name)
		}
		tasks[i] = task
	}
	return tasks, nil
}
