package devtrack

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devtrack",
	Short: "DevTrack - developer time tracking tool",
	Long: `DevTrack to track dev's activity`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
