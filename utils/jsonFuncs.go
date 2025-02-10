package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (t Task) String() string {
	return fmt.Sprintf("Id: %v\nDescription: %v\nStatus: %v\nCreatedAt: %v", t.Id, t.Description, t.Status, t.CreatedAt)
}

func LoadTasks(fileName string) (file *os.File, tasks []Task) {
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatalf("Error openning file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	return
}

func SaveTasks(fileName string, tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}
