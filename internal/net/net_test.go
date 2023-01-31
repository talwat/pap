package net_test

import (
	"testing"

	"github.com/talwat/pap/internal/net"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type Todo struct {
		ID int
	}

	var todo Todo

	net.Get("https://jsonplaceholder.typicode.com/todos/1", "todo not found", &todo)

	if todo.ID != 1 {
		t.Errorf(`Get("https://jsonplaceholder.typicode.com/todos/1", &todo) = %d; want 1`, todo.ID)
	}
}
