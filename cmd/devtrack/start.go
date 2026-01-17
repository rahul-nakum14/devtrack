package devtrack

import (
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking a new work",
	Long:  "Start a new work session and begin trace work..",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		fmt.Println("Started tracking:", task)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
