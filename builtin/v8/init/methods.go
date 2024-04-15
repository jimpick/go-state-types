package init

import (
	"github.com/jimpick/go-state-types/abi"
	"github.com/jimpick/go-state-types/builtin"
)

var Methods = map[abi.MethodNum]builtin.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) *abi.EmptyValue)}, // Constructor
	2: {"Exec", *new(func(*ExecParams) *ExecReturn)},                   // Exec
}
