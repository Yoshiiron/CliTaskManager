package cmd

import (
	"CliTaskManager/internal"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "Lists the issues.",
	Run: func(cmd *cobra.Command, args []string) {
		file, tasks := internal.LoadTasks()
		defer file.Close()

		if len(args) == 0 {
			fmt.Printf("%-3s | %-30s | %-10s\n", "ID", "Description", "Status")
			fmt.Println("---------------------------------------------")
			for _, task := range tasks {
				fmt.Printf("%-3d | %-30s | %-10s\n", task.Id, task.Description, task.Status)
			}
		} else if strings.ToLower(args[0]) == "done" {
			fmt.Printf("%-3s | %-30s\n", "ID", "Description")
			fmt.Println("---------------------------------------------")
			for _, task := range tasks {
				if task.Status == "Done" {
					fmt.Printf("%-3d | %-30s\n", task.Id, task.Description)
				}
			}
		} else if strings.ToLower(args[0]) == "todo" {
			fmt.Printf("%-3s | %-30s\n", "ID", "Description")
			fmt.Println("---------------------------------------------")
			for _, task := range tasks {
				if task.Status == "ToDo" {
					fmt.Printf("%-3d | %-30s\n", task.Id, task.Description)
				}
			}
		} else if strings.ToLower(args[0]) == "in-progress" {
			fmt.Printf("%-3s | %-30s\n", "ID", "Description")
			fmt.Println("---------------------------------------------")
			for _, task := range tasks {
				if task.Status == "InProgress" {
					fmt.Printf("%-3d | %-30s\n", task.Id, task.Description)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
