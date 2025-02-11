package cmd

import (
	"CliTaskManager/utils"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Mark the task as Done.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("error: not enough args")
			return
		}

		file, tasks := utils.LoadTasks()
		defer file.Close()

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println("error: first argument must be an task ID.")
			return
		}

		for ind, task := range tasks {
			if task.Id == taskID {
				if task.Status == "Done" {
					fmt.Println("This issue is already done.")
					continue
				}
				tasks[ind].Status = "Done"
				tasks[ind].UpdatedAt = time.Now().Format(time.RFC822)
				fmt.Println(tasks[ind])
			} else {
			}
		}

		utils.SaveTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(markDoneCmd)
}
