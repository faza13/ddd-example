package domain

import "context"

type ProgramStudiRepository interface {
	List(ctx context.Context) ([]ProgramStudi, error)
}
