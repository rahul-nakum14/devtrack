package devtrack

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	 "github.com/rahul-nakum14/devtrack/internal/db"
	"github.com/rahul-nakum14/devtrack/internal/repository"
	"github.com/rahul-nakum14/devtrack/internal/service"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "show statistic",
}

var statsTodayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today stats",
	RunE:  statsToday,
}

func statsToday(cmd *cobra.Command, args []string) error {
	database, err := db.OpenDB()
	if err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	sessionService := service.NewSessionService(repo)

	perTask, total, err := sessionService.GetTodayStats()
	if err != nil {
		return err
	}

	fmt.Println("Stats for today:---------->")

	for task, duration := range perTask {
		fmt.Printf("%-10s %s\n", task, duration.Round(time.Minute))
	}

	fmt.Println("----------------")
	fmt.Println("Total     ", total.Round(time.Minute))

	return nil
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.AddCommand(statsTodayCmd)
}
