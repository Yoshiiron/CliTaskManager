package cmd

import (
	"CliTaskManager/internal"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// addtaskCmd represents the addtask command
var addtaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add's task.",
	Long:  `This command will add's task.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := 1

		if len(args) == 0 {
			log.Println("error: description is empty")
			return
		}

		file, tasks := internal.LoadTasks()
		defer file.Close()

		if len(tasks) != 0 {
			id = tasks[len(tasks)-1].Id + 1
		}

		task.Id = id
		task.Description = args[0]
		task.Status = "ToDo"
		task.CreatedAt = time.Now().Format(time.RFC822)
		task.UpdatedAt = "Not updated yet."
		tasks = append(tasks, task)

		internal.SaveTasks(tasks)

		fmt.Println(task)
	},
}

var (
	task internal.Task
)

func init() {
	rootCmd.AddCommand(addtaskCmd)
}
