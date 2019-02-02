package main

import (
	"context"
	"fmt"

	"github.com/junzexu/nuts/demo/gen-go/multiple"
)

// MultiplicationServiceHandler .
type MultiplicationServiceHandler struct {
}

// Multiply .
func (h *MultiplicationServiceHandler) Multiply(ctx context.Context, n1 multiple.Int, n2 multiple.Int) (multiple.Int, error) {
	fmt.Printf("called: %+v, %+v => result: %+v\n", n1, n2, n1*n2)
	return n1 * n2, nil
}
