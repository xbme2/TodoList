package service


/*
	task 在终端的展示
*/
import (
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func ShowTask() {
	tasks := ReadCsv("test.csv")
	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.SetPageSize(10)
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name: "任务名称",
			ColorsHeader: text.Colors{
				text.BgCyan,
			},
		},
		{
			Name: "优先程度",
			Transformer: func(val interface{}) string {
				if val == "1" {
					return text.Colors{text.FgBlue}.Sprintf("Normal")
				} else if val == "0" {
					return text.Colors{text.FgWhite}.Sprintf("Light")
				} else if val == "2" {
					return text.Colors{text.FgHiYellow}.Sprintf("Important")
				} else if val == "3" {
					return text.Colors{text.FgHiRed}.Sprintf("Critical")
				} else {
					return ""
				}
			},
		},
	})

	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"任务名称", "截止日期", "优先程度"})
	for _, arr := range tasks {
		t.AppendRow(table.Row{arr[1], arr[2], arr[3]})
	}
	t.Render()

}
