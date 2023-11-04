package domain

import "context"

type StudentRepository interface {
	ByID(ctx context.Context, id uint64) (Student, error)
}
