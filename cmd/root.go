package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todoList ([-c | -d]) ",
	Short: "Welcome to xbme2's todoList!",
	Long: `A Simole and Simple todoList with Cobra in Go, 
		   By xbme2
		   You can click '' for detail`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(cmd.Short)

	},
}

func Execute() {
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(showCommand)
	addCommand.Flags().StringArrayP("name", "n", []string{}, "任务主题") // 添加短标志 -s
	addCommand.Flags().StringArrayP("deadline", "d", []string{}, "截止日期")
	addCommand.Flags().StringArrayP("level", "l", []string{}, "关键程度")
	// addCommand.Flags().Str
	// rootCmd.AddCommand(ver)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
