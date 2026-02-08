package devtrack

import (
	"fmt"
	
	"github.com/rahul-nakum14/devtrack/internal/repository"
	"github.com/rahul-nakum14/devtrack/internal/db"
	"github.com/rahul-nakum14/devtrack/internal/service"
	"github.com/spf13/cobra"
)


var stopCmd = &cobra.Command{
	Use: "stop",
	Short: "stop task",
	Long:"stop Exisitng task",
	RunE: stopTracking,
}

func stopTracking(cmd *cobra.Command, args []string) error {
	database, err := db.OpenDB()
	if  err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	sessionService := service.NewSessionService(repo)

	session, err := sessionService.StopSession()
	if err != nil {
		return err
	}

	fmt.Println("Stopped tracking:", session.Task)
	return nil
}

func init() {
    rootCmd.AddCommand(stopCmd)
}