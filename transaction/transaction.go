package transaction

import (
	"context"
)

type Transaction interface {
	DoInTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}

type Noop struct {
}

func (n *Noop) DoInTx(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	return f(ctx)
}
