package task

import "testing"

type TestData struct {
	name        string
	description string
	priority    int
	done        bool
	position    int
	expected    string
}

func TestTask_PrettyPriority(t *testing.T) {
	tests := []TestData{
		{
			name:     "high-priority",
			priority: 1,
			expected: "[HIGH]",
		},
		{
			name:     "medium-priority",
			priority: 2,
			expected: "[MEDIUM]",
		},
		{
			name:     "low-priority",
			priority: 3,
			expected: "[LOW]",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{
				Description: test.description,
				Priority:    test.priority,
			}
			if got := task.PrettyPriority(); got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}

func TestTask_PrettyStatus(t *testing.T) {
	tests := []TestData{
		{
			name:     "done-status",
			priority: 1,
			done:     true,
			expected: "(DONE)",
		},
		{
			name:     "todo-status",
			priority: 2,
			done:     false,
			expected: "(TODO)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{
				Description: test.description,
				Done:        test.done,
			}
			if got := task.PrettyStatus(); got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}

func TestTask_PrettyPosition(t *testing.T) {
	tests := []TestData{
		{
			name:     "1",
			position: 1,
			expected: "1.",
		},
		{
			name:     "123",
			position: 123,
			expected: "123.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{
				Description: test.description,
				position:    test.position,
			}
			if got := task.PrettyPosition(); got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}
