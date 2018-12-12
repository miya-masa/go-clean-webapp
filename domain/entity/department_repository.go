package entity

import "context"

type DepartmentRepository interface {
	Find(ctx context.Context, id string) (*Department, error)
	Store(ctx context.Context, department *Department) error
	Delete(ctx context.Context, id string) (int, error)
}
