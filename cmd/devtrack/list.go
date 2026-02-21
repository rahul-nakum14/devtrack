package devtrack

import (
	"fmt"
	"time"
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

		fmt.Printf("ID: %d | Task: %s | Project: %s\n",
			s.ID, s.Task, s.Project)
	
		start := s.StartTime.Local().Format("02 Jan 2006 15:04")
		fmt.Printf("Start: %s\n", start)
	
		var duration time.Duration
	
		if s.EndTime == nil {
			fmt.Println("End:   Running")
			duration = time.Since(s.StartTime)
		} else {
			end := s.EndTime.Local().Format("02 Jan 2006 15:04")
			fmt.Printf("End:   %s\n", end)
			duration = s.EndTime.Sub(s.StartTime)
		}
	
		min := int(duration.Minutes())
		sec := int(duration.Seconds()) % 60
	
		fmt.Printf("Spent: %dm %ds\n", min, sec)
		fmt.Println("-----------------------------")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listAllSessionsCmd)
}