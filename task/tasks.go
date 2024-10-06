package task

import (
	"encoding/json"
	"errors"
	"os"
)

type Tasks []Task

func (t Tasks) DeleteTask(taskID int) (Tasks, error) {
	if taskID <= 0 {
		return []Task{}, errors.New("task-id can not be negative")
	}
	if taskID > len(t) {
		return []Task{}, errors.New("task-id does not exist")
	}
	taskID -= 1
	t = append(t[:taskID], t[taskID+1:]...)
	return t, nil
}

func (t Tasks) FinishTask(taskID int) (Tasks, error) {
	if taskID <= 0 {
		return []Task{}, errors.New("task-id can not be negative")
	}
	if taskID > len(t) {
		return []Task{}, errors.New("task-id does not exist")
	}
	taskID -= 1
	t[taskID].SetDone()
	return t, nil
}

func (t Tasks) UpdateTaskDescription(taskID int, description string) (Tasks, error) {
	if taskID <= 0 {
		return []Task{}, errors.New("task-id can not be negative")
	}
	if taskID > len(t) {
		return []Task{}, errors.New("task-id does not exist")
	}
	if description == "" {
		return []Task{}, errors.New("task description can not be blank")
	}
	taskID -= 1
	t[taskID].SetDescription(description)
	return t, nil
}

func ReadTasks(filename string) (Tasks, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return []Task{}, err
	}
	for i, _ := range tasks {
		tasks[i].position = i + 1
	}
	return tasks, nil
}

func SaveTasks(filename string, tasks Tasks) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

type ByPriority Tasks

func (s ByPriority) Len() int {
	return len(s)
}

func (s ByPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPriority) Less(i, j int) bool {
	if s[i].Done == s[j].Done {
		if s[i].Priority == s[j].Priority {
			return s[i].position < s[j].position
		}
		return s[i].Priority < s[j].Priority
	}
	return !s[i].Done
}
