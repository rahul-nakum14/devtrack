package devtrack

import (
	"github.com/spf13/cobra"

	"devtrack/internal/service"
)

var startCmd = &cobra.Command{
	Use:   "start [task]",
	Short: "Start tracking a new work",
	Long:  "Start a new work session and begin tracking work.",
	Args:  cobra.ExactArgs(1),  // here it validates the args like devtrack start <1 args must here>
	Run:   startTracking,
}

func startTracking(cmd *cobra.Command, args []string) {
	task := args[0]
	project, _ := cmd.Flags().GetString("project")

	//coonect with db and stored sessions


	database, err := db.OpenDB()
	if err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	sessionService := service.NewSessionService(repo)

	session, err := sessionService.StartSession(task, project)
	if err != nil {
		return err
	}

	// sessionService := service.NewSessionService()
	// sessionService.StartSession(task, project)
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringP(
		"project",
		"p",
		"",
		"Project name",
	)
}
