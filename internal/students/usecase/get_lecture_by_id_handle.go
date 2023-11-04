package usecase

import (
	"app/internal/students/domain"
	"context"
)

type getByIDHandle struct {
	studentRepository domain.StudentRepository
}

type GetByIDHandle interface {
	Handle(ctx context.Context, ID uint64) (domain.Student, error)
}

func NewGetByIDHandle(repository domain.StudentRepository) GetByIDHandle {
	return getByIDHandle{studentRepository: repository}
}

func (l getByIDHandle) Handle(ctx context.Context, id uint64) (domain.Student, error) {
	return l.studentRepository.ByID(ctx, id)
}
