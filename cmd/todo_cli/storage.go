package todocli

import (
	"os"
	"strings"
)

const taskFile = "tasks.txt"

// Load tasks from file into Todo_list
func loadTasks() error {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			Todo_list = []string{} // no file yet
			return nil
		}
		return err
	}
	content := strings.TrimSpace(string(data))
	if content == "" {
		Todo_list = []string{}
	} else {
		Todo_list = strings.Split(content, "\n")
	}
	return nil
}

// Save Todo_list to file
func saveTasks() error {
	return os.WriteFile(taskFile, []byte(strings.Join(Todo_list, "\n")), 0644)
}
