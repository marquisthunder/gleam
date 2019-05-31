package executor

import (
	"github.com/marquisthunder/gleam/flow"
	"github.com/marquisthunder/gleam/sql/expression"
)

type Executor interface {
	Exec() *flow.Dataset
	Schema() expression.Schema
}
