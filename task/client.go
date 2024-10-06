package task

import (
	"encoding/json"
	"errors"
	"os"
)

type Client struct {
	tasks []Task
}

func (c *Client) AddTask(task Task) {
	c.tasks = append(c.tasks, task)
}

func (c *Client) CleanTasks() {
	c.tasks = nil
}

func (c *Client) DeleteTask(taskID int) error {
	if taskID <= 0 {
		return errors.New("task-id can not be negative")
	}
	if taskID > len(c.tasks) {
		return errors.New("task-id does not exist")
	}
	taskID -= 1
	c.tasks = append(c.tasks[:taskID], c.tasks[taskID+1:]...)
	return nil
}

func (c *Client) FinishTask(taskID int) error {
	if taskID <= 0 {
		return errors.New("task-id can not be negative")
	}
	if taskID > len(c.tasks) {
		return errors.New("task-id does not exist")
	}
	taskID -= 1
	c.tasks[taskID].SetDone()
	return nil
}

func (c *Client) UpdateTaskDescription(taskID int, description string) error {
	if taskID <= 0 {
		return errors.New("task-id can not be negative")
	}
	if taskID > len(c.tasks) {
		return errors.New("task-id does not exist")
	}
	if description == "" {
		return errors.New("task description can not be blank")
	}
	taskID -= 1
	c.tasks[taskID].SetDescription(description)
	return nil
}

func (c *Client) ReadTasks(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.tasks)
	if err != nil {
		return err
	}
	for i := range c.tasks {
		c.tasks[i].position = i + 1
	}
	return nil
}

func (c *Client) SaveTasks(filename string) error {
	data, err := json.Marshal(c.tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

type ByPriority Client

func (s ByPriority) Len() int {
	return len(s.tasks)
}

func (s ByPriority) Swap(i, j int) {
	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
}

func (s ByPriority) Less(i, j int) bool {
	if s.tasks[i].Done == s.tasks[j].Done {
		if s.tasks[i].Priority == s.tasks[j].Priority {
			return s.tasks[i].position < s.tasks[j].position
		}
		return s.tasks[i].Priority < s.tasks[j].Priority
	}
	return !s.tasks[i].Done
}
