package repo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDummyGetTodoList(t *testing.T) {
	repo := NewDummyTodoRepo()
	todoList, err := repo.getTodoList(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, pageSize, len(todoList))
	assert.Equal(t, "1", todoList[0].ID)
	assert.Equal(t, "2", todoList[1].ID)
	assert.Equal(t, "3", todoList[2].ID)

	todoList, err = repo.getTodoList(context.Background(), "4")
	assert.Nil(t, err)
	assert.Equal(t, pageSize, len(todoList))
	assert.Equal(t, "5", todoList[0].ID)
	assert.Equal(t, "6", todoList[1].ID)
	assert.Equal(t, "7", todoList[2].ID)

	todoList, err = repo.getTodoList(context.Background(), "9")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(todoList))
	assert.Equal(t, "10", todoList[0].ID)
}

func TestDummyGetTodo(t *testing.T) {
	repo := NewDummyTodoRepo()
	todo, err := repo.getTodo(context.Background(), "1")
	assert.Nil(t, err)
	assert.Equal(t, "1", todo.ID)

	todo, err = repo.getTodo(context.Background(), "5")
	assert.Nil(t, err)
	assert.Equal(t, "5", todo.ID)

	todo, err = repo.getTodo(context.Background(), "10")
	assert.Nil(t, err)
	assert.Equal(t, "10", todo.ID)

	todo, err = repo.getTodo(context.Background(), "11")
	assert.NotNil(t, err)
}
