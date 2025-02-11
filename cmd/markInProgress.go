package cmd

import (
	"CliTaskManager/internal"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// markInProgressCmd represents the markInProgress command
var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Mark the task as In-Progress.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("error: not enough args")
			return
		}

		file, tasks := internal.LoadTasks()
		defer file.Close()

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println("error: first argument must be an task ID.")
			return
		}

		for ind, task := range tasks {
			if task.Id == taskID {
				if task.Status == "InProgress" {
					fmt.Println("This issue is already InProgress.")
					fmt.Println(tasks[ind])
					break
				}
				tasks[ind].Status = "InProgress"
				tasks[ind].UpdatedAt = time.Now().Format(time.RFC822)
				fmt.Println(tasks[ind])
			}
		}

		internal.SaveTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}
