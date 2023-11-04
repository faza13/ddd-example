package lectures

import (
	"app/internal/lectures/interfaces/private/intraprocess"
	"app/internal/user/domain"
	"context"
)

type intraprocessService struct {
	intraprocessInterface intraprocess.LectureInterface
}

func NewIntraprocessService(intraprocessInterface intraprocess.LectureInterface) intraprocessService {
	return intraprocessService{intraprocessInterface}
}

func (i intraprocessService) LectureByID(ctx context.Context, id uint64) (domain.Lecture, error) {
	domainLecture, err := i.intraprocessInterface.LectureByID(ctx, id)
	if err != nil {
		return domain.Lecture{}, err
	}

	return domain.Lecture{
		ID:           domainLecture.ID,
		Title:        domainLecture.Title,
		Gelar:        domainLecture.Gelar,
		Nama:         domainLecture.Nama,
		Kelamin:      domainLecture.Kelamin,
		Telepon:      domainLecture.Telepon,
		HP:           domainLecture.HP,
		Email:        domainLecture.Email,
		TanggalLahir: domainLecture.TanggalLahir,
	}, err
}
