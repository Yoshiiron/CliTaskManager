package main

import (
	"CliTaskManager/cmd"
	"CliTaskManager/utils"
)

func main() {
	file, _ := utils.LoadTasks("tasks.json")
	defer file.Close()
	cmd.Execute()
}
