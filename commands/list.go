package commands

import (
	"drach/db"
	"drach/models"
	"fmt"
)

func ListCmd() {
	tasks, err := models.ListTasks(db.DB)
	if err != nil {
		fmt.Printf("Error on list tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No task found")
		return
	}

	fmt.Println("Tasks: ")
	for _, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d, %s [%s] - Created in: %s\n",
			task.ID,
			task.Description,
			status,
			task.CreatedAt.Format("02/01/2006 15:04"))
	}
}
