/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"CliTaskManager/internal"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete task.",
	Long:  `Delete task from todo app.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("error: id needed")
			return
		}

		file, tasks := internal.LoadTasks()
		defer file.Close()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
			return
		}

		for index, task := range tasks {
			if task.Id == id {
				tasks = append(tasks[:index], tasks[index+1:]...)
			}
		}

		internal.SaveTasks(tasks)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
