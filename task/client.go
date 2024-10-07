package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Tasks []Task

type Client struct {
	tasks Tasks
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

func (c *Client) DisplayTasks(sortPriority bool, displayPriority bool) {
	if len(c.tasks) == 0 {
		fmt.Println("Nothing on your To-Do list for the moment.")
	} else {
		w := tabwriter.NewWriter(os.Stdout, 4, 0, 1, ' ', 0)
		if displayPriority {
			fmt.Fprintln(w, "ID\tTask\tStatus\tPriority")
			fmt.Fprintln(w, "--\t----\t------\t--------")
		} else {
			fmt.Fprintln(w, "ID\tTask\tStatus")
			fmt.Fprintln(w, "--\t----\t------")
		}
		if sortPriority {
			sort.Sort(ByPriority(c.tasks))
		}
		for _, t := range c.tasks {
			if displayPriority {
				fmt.Fprintln(w, t.PrettyPosition()+"\t"+t.Description+"\t"+t.PrettyStatus()+"\t"+t.PrettyPriority())
			} else {
				fmt.Fprintln(w, t.PrettyPosition()+"\t"+t.Description+"\t"+t.PrettyStatus())
			}
		}
		w.Flush()
	}
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
