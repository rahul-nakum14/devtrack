package devtrack

import (
	"fmt"
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

	fmt.Println("Stats for today:")
	
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
    rootCmd.AddCommand(statsCmd)
    statsCmd.AddCommand(statsTodayCmd)
}
