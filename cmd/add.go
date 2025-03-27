package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"main.go/global"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: global.Iswroking,
	Long:  "thc command is used to add one or more tasks into your list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
	},
}
