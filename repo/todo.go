package repo

import (
	"context"
	"fmt"
	"github.com/pistatium/graphql_study/graph/model"
	"strconv"
)

type TodoRepository interface {
	getTodoList(ctx context.Context, cursor string) ([]*model.Todo, error)
	getTodo(ctx context.Context, id string) (*model.Todo, error)
}

const pageSize = 3

type DummyTodoRepo struct {
	todoList []*model.Todo
}

func (d DummyTodoRepo) getTodoList(ctx context.Context, cursor string) ([]*model.Todo, error) {
	resp := make([]*model.Todo, 0)
	for _, todo := range d.todoList {
		if cursor != "" {
			cursorID, err := strconv.Atoi(cursor)
			if err != nil {
				return nil, fmt.Errorf("cursor is not a number: %s", cursor)
			}
			todoID, err := strconv.Atoi(todo.ID)
			if err != nil {
				return nil, fmt.Errorf("todo id is not a number %s", todo.ID)
			}
			if cursorID >= todoID {
				continue
			}
		}
		resp = append(resp, todo)
		if len(resp) >= pageSize {
			break
		}
	}
	return resp, nil
}

func (d DummyTodoRepo) getTodo(ctx context.Context, id string) (*model.Todo, error) {
	for _, todo := range d.todoList {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func NewDummyTodoRepo() TodoRepository {
	user1 := model.User{
		ID:   "1",
		Name: "hoge",
	}
	user2 := model.User{
		ID:   "2",
		Name: "fuga",
	}
	return DummyTodoRepo{
		todoList: []*model.Todo{
			{
				ID:   "1",
				Text: "todo 1",
				Done: false,
				User: &user1,
			},
			{
				ID:   "2",
				Text: "todo 2",
				Done: false,
				User: &user1,
			},
			{
				ID:   "3",
				Text: "todo 3",
				Done: true,
				User: &user1,
			},
			{
				ID:   "4",
				Text: "todo 4",
				Done: false,
				User: &user2,
			},
			{
				ID:   "5",
				Text: "todo 5",
				Done: false,
				User: &user2,
			},
			{
				ID:   "6",
				Text: "todo 6",
				Done: false,
				User: &user2,
			}, {
				ID:   "7",
				Text: "todo 7",
				Done: false,
				User: &user2,
			},
			{
				ID:   "8",
				Text: "todo 8",
				Done: false,
				User: &user1,
			},
			{
				ID:   "9",
				Text: "todo 9",
				Done: true,
				User: &user2,
			},
			{
				ID:   "10",
				Text: "todo 10",
				Done: false,
				User: &user1,
			},
		},
	}
}
