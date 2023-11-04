package usecase

import (
	"app/internal/lectures/domain"
	"context"
)

type getByIDHandle struct {
	lectureRepository domain.LectureRepository
}

type GetByIDHandle interface {
	Handle(ctx context.Context, ID uint64) (domain.Lecture, error)
}

func NewGetByIDHandle(repository domain.LectureRepository) GetByIDHandle {
	return getByIDHandle{lectureRepository: repository}
}

func (l getByIDHandle) Handle(ctx context.Context, id uint64) (domain.Lecture, error) {
	return l.lectureRepository.GetByID(ctx, id)
}
