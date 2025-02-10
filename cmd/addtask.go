package cmd

import (
	"CliTaskManager/utils"
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

		file, tasks := utils.LoadTasks("tasks.json")
		defer file.Close()

		if len(tasks) != 0 {
			id = tasks[len(tasks)-1].Id + 1
		}

		task.Id = id
		task.Description = args[0]
		task.Status = "todo"
		task.CreatedAt = time.Now().Format(time.RFC822)
		tasks = append(tasks, task)

		utils.SaveTasks("tasks.json", tasks)

		fmt.Println(task)
	},
}

var (
	task utils.Task
)

func init() {
	rootCmd.AddCommand(addtaskCmd)
}
