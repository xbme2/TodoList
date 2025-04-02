package cmd

import (
	"github.com/spf13/cobra"
	"main.go/service"
)

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "show task list into the cmd",
	Long:  "show task list into the cmd;show task list into the cmd",
	Run: func(cmd *cobra.Command, args []string) {
		service.ShowTask()
	},
}
