package devtrack

import (
	"fmt"

	"github.com/rahul-nakum14/devtrack/internal/db"
	"github.com/rahul-nakum14/devtrack/internal/repository"
	"github.com/rahul-nakum14/devtrack/internal/service"
	"github.com/spf13/cobra"

)

var statsWeekCmd = &cobra.Command{
	Use:   "week",
	Short: "Show weekly stats",
	RunE:  statsWeek,
}

func statsWeek(cmd *cobra.Command, args []string) error {

	database, err := db.OpenDB()
	if err != nil {
		return err
	}

	repo := repository.NewSessionSQLiteRepository(database)
	sessionService := service.NewSessionService(repo)

	perTask, total, err := sessionService.GetWeekStats()
	if err != nil {
		return err
	}

	 fmt.Println("Stats for last week day:")

	for task, duration := range perTask {
		min := int(duration.Minutes())
		sec := int(duration.Seconds()) % 60
		fmt.Printf("%-10s %dm %ds\n", task, min, sec)
	}

	fmt.Println("----------------")

	totalMin := int(total.Minutes())
	totalSec := int(total.Seconds()) % 60
	fmt.Printf("Total      %dm %ds\n", totalMin, totalSec)

	return nil
}

func init() {
	statsCmd.AddCommand(statsWeekCmd)
}
