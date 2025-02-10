/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"CliTaskManager/utils"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [taskID] [description]",
	Short: "Update task description.",
	Long:  `Update task description.`,
	Run: func(cmd *cobra.Command, args []string) {

		file, tasks := utils.LoadTasks()
		defer file.Close()

		if len(args) < 2 {
			log.Println("error: not enough args to complete the command.")
			return
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println("error: first argument must be an integer task ID.")
			return
		}

		description := args[1]
		if description == "" {
			log.Println("error: description cannot be empty or must be int.")
			return
		}

		for ind, task := range tasks {
			if task.Id == taskID {
				tasks[ind].Description = args[1]
				tasks[ind].UpdatedAt = time.Now().Format(time.RFC822)
				fmt.Println(tasks[ind])
			}
		}

		utils.SaveTasks(tasks)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
