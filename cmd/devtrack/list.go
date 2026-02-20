package devtrack

import (
	"fmt"
	// "time"

	"github.com/rahul-nakum14/devtrack/internal/db"
	"github.com/rahul-nakum14/devtrack/internal/repository"
	"github.com/rahul-nakum14/devtrack/internal/service"
	"github.com/spf13/cobra"
)

var listAllSessionsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sessions",
	RunE:  listSessions,
}

func listSessions(cmd *cobra.Command, args []string) error {

	database, err := db.OpenDB()
	if err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	service := service.NewSessionService(repo)

	sessions, err := service.GetAllSessions()
	if err != nil {
		return err
	}

	for _, s := range sessions {
		fmt.Printf("ID: %d\n", s.ID)
		fmt.Printf("Task: %s\n", s.Task)
		fmt.Printf("Project: %s\n", s.Project)
		fmt.Printf("Start: %v\n", s.StartTime)
		fmt.Printf("End: %v\n", s.EndTime)
		fmt.Println("-----------")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listAllSessionsCmd)
}