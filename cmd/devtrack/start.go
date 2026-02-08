package devtrack

import (
	"fmt"
	
	"github.com/rahul-nakum14/devtrack/internal/repository"
	"github.com/rahul-nakum14/devtrack/internal/db"
	"github.com/rahul-nakum14/devtrack/internal/service"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [task]",
	Short: "Start tracking a new work",
	Long:  "Start a new work session and begin tracking work.",
	Args:  cobra.ExactArgs(1), 
	RunE:   startTracking,
}

func startTracking(cmd *cobra.Command, args []string) error {
	task := args[0]
	project, _ := cmd.Flags().GetString("project")

	//coonect with db and stored sessions

	database, err := db.OpenDB()
	if err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	
	sessionService := service.NewSessionService(repo)

	if err := repo.Migrate(); err != nil {
		return err
	}
	session, err := sessionService.StartSession(task, project)
	if err != nil {
		return err
	}

	fmt.Println("Started tracking:", session.Task)
	return nil
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
