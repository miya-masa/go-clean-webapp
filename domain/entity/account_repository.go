package entity

import "context"

type AccountRepository interface {
	List(ctx context.Context) ([]*Account, error)
	Find(ctx context.Context, id string) (*Account, error)
	Store(ctx context.Context, account *Account) (*Account, error)
	Delete(ctx context.Context, id string) (int, error)
}
