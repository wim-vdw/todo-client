package task

import (
	"strconv"
)

type Task struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Done        bool   `json:"done"`
	position    int
}

func (t *Task) SetPriority(priority int) {
	t.Priority = priority
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
