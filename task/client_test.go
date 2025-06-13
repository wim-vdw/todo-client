package task

import (
	"os"
	"testing"
)

func TestClient_AddTask(t *testing.T) {
	client := &Client{}
	task := Task{Description: "Test Task", Priority: 2}
	client.AddTask(task)

	if len(client.tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(client.tasks))
	}
	if client.tasks[0].Description != "Test Task" {
		t.Errorf("expected task description 'Test Task', got '%s'", client.tasks[0].Description)
	}
}

func TestClient_DeleteTask(t *testing.T) {
	client := &Client{tasks: Tasks{
		{Description: "Task 1"},
		{Description: "Task 2"},
	}}
	err := client.DeleteTask(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(client.tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(client.tasks))
	}
	if client.tasks[0].Description != "Task 2" {
		t.Errorf("expected remaining task 'Task 2', got '%s'", client.tasks[0].Description)
	}
}

func TestClient_FinishTask(t *testing.T) {
	client := &Client{tasks: Tasks{
		{Description: "Task 1", Done: false},
	}}
	err := client.FinishTask(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !client.tasks[0].Done {
		t.Errorf("expected task to be marked as done")
	}
}

func TestClient_UpdateTaskDescription(t *testing.T) {
	client := &Client{tasks: Tasks{
		{Description: "Old Description"},
	}}
	err := client.UpdateTaskDescription(1, "New Description")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if client.tasks[0].Description != "New Description" {
		t.Errorf("expected updated description 'New Description', got '%s'", client.tasks[0].Description)
	}
}

func TestClient_UpdateTaskPriority(t *testing.T) {
	client := &Client{tasks: Tasks{
		{Priority: 1},
	}}
	err := client.UpdateTaskPriority(1, 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if client.tasks[0].Priority != 3 {
		t.Errorf("expected updated priority 3, got %d", client.tasks[0].Priority)
	}
}

func TestClient_DisplayTasks(t *testing.T) {
	client := &Client{tasks: Tasks{
		{Description: "Task 1", Priority: 1, Done: false},
		{Description: "Task 2", Priority: 2, Done: true},
	}}
	client.DisplayTasks(true, true)
	// ToDo: This test is visual; ensure the output matches expectations manually.
}

func TestClient_ReadTasks(t *testing.T) {
	filename := "test_tasks.json"
	data := `[{"Description":"Task 1","Priority":1,"Done":false},{"Description":"Task 2","Priority":2,"Done":true}]`
	os.WriteFile(filename, []byte(data), 0644)
	defer os.Remove(filename)

	client := &Client{Filename: filename}
	err := client.ReadTasks()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(client.tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(client.tasks))
	}
}

func TestClient_SaveTasks(t *testing.T) {
	filename := "test_tasks.json"
	defer os.Remove(filename)

	client := &Client{Filename: filename, tasks: Tasks{
		{Description: "Task 1", Priority: 1, Done: false},
		{Description: "Task 2", Priority: 2, Done: true},
	}}
	err := client.SaveTasks()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	data, _ := os.ReadFile(filename)
	expected := `[{"description":"Task 1","priority":1,"done":false},{"description":"Task 2","priority":2,"done":true}]`
	if string(data) != expected {
		t.Errorf("expected saved data '%s', got '%s'", expected, string(data))
	}
}
