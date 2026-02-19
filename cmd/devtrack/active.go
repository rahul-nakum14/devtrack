package devtrack

import (
	"fmt"
	
	"github.com/rahul-nakum14/devtrack/internal/repository"
	"github.com/rahul-nakum14/devtrack/internal/db"
	"github.com/rahul-nakum14/devtrack/internal/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "active",
	Short: "List active sessions",
	Long:  "List active sessions details",
	RunE:   listAciveSessions,
}

func listAciveSessions(cmd *cobra.Command, args []string) error {

	//coonect with db and stored sessions

	database, err := db.OpenDB()
	if err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	
	sessionService := service.NewSessionService(repo)
	
	session, err := sessionService.GetActiveSession()
	if err != nil {
		return err
	}

	fmt.Println("Seesions Are", session.Task)
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	// startCmd.Flags().StringP(
	// 	"project",
	// 	"p",
	// 	"",
	// 	"Project name",
	// )
}
