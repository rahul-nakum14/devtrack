package devtrack

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking a new work",
	Long:  "Start a new work session and begin tracking",
	Args:  cobra.ExactArgs(1),
	Run:   startTracking,
}

func startTracking(cmd *cobra.Command, args []string) {
	task := args[0]

	project, _ := cmd.Flags().GetString("project")

	fmt.Println("Started tracking:", task)

	if project != "" {
		fmt.Println("Project:", project)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP(
		"project", // flag name
		"p",       // shorthand
		"",        // default value
		"Project name", //descrtiption
	)
}
