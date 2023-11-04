package usecase

import (
	"app/internal/program_studi/domain"
	"context"
)

type listHandle struct {
	programStudyRepository domain.ProgramStudiRepository
}
type ListHandle interface {
	Handle(ctx context.Context) ([]domain.ProgramStudi, error)
}

func NewListHandle(repository domain.ProgramStudiRepository) ListHandle {
	return listHandle{programStudyRepository: repository}
}

func (h listHandle) Handle(ctx context.Context) ([]domain.ProgramStudi, error) {
	return h.programStudyRepository.List(ctx)
}
