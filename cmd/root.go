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

		fmt.Println("haha")

	},
}

func Execute() {
	rootCmd.AddCommand(addCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
