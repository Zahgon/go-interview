package postfixes

import (
	"errors"

	"github.com/shomali11/go-interview/datastructures/stacks/slicestacks"
)

var (
	errERR = errors.New("#ERR")
)

// Evaluate evaluates postfix expressions
func Evaluate(expression string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func isNumber(text string) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

func pop(stack *slicestacks.Stack[float64]) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
