package domain

import "context"

type LectureRepository interface {
	GetByID(ctx context.Context, id uint64) (Lecture, error)
}
