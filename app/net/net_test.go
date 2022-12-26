package net

import "testing"

func TestGet(t *testing.T) {
	type Todo struct {
		ID int
	}

	var todo Todo

	Get("https://jsonplaceholder.typicode.com/todos/1", &todo)

	if todo.ID != 1 {
		t.Errorf(`Get("https://jsonplaceholder.typicode.com/todos/1", &todo) = %d; want 1`, todo.ID)
	}
}
