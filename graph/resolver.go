package graph

import (
	"github.com/pistatium/graphql_study/repo"
)

type Resolver struct {
	TodoRepo repo.TodoRepository
}
