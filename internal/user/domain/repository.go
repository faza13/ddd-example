package domain

import "context"

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (User, error)
}
