package task

import (
	"errors"
	"strconv"
)

const (
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorReset  = "\u001b[0m"
)

type Task struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Done        bool   `json:"done"`
	position    int
}

func (t *Task) SetPriority(priority int) error {
	if priority != 1 && priority != 2 && priority != 3 {
		return errors.New("invalid priority: must be 1, 2, or 3")
	}
	t.Priority = priority
	return nil
}

func (t *Task) SetDone() {
	t.Done = true
}

func (t *Task) SetDescription(description string) {
	t.Description = description
}

func (t *Task) PrettyPriority() string {
	switch t.Priority {
	case 1:
		return "[HIGH]"
	case 3:
		return "[LOW]"
	default:
		return "[MEDIUM]"
	}
}

func (t *Task) PrettyPosition() string {
	return strconv.Itoa(t.position) + "."
}

func (t *Task) PrettyStatus() string {
	if t.Done {
		return "(DONE)"
	} else {
		return "(TODO)"
	}
}

func (t *Task) PrettyColorStatus() string {
	if t.Done {
		return ColorGreen + "(DONE)" + ColorReset
	} else {
		return ColorYellow + "(TODO)" + ColorReset
	}
}
