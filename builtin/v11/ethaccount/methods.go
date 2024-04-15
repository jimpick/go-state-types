package ethaccount

import (
	"github.com/jimpick/go-state-types/abi"
	"github.com/jimpick/go-state-types/builtin"
)

var Methods = map[abi.MethodNum]builtin.MethodMeta{
	1: {"Constructor", *new(func(value *abi.EmptyValue) *abi.EmptyValue)},
}
