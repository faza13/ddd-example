package lectures

import (
	"app/internal/students/interfaces/private/intraprocess"
	"app/internal/user/domain"
	"context"
)

type intraprocessService struct {
	intraprocessInterface intraprocess.StudentInterface
}

func NewIntraprocessService(intraprocessInterface intraprocess.StudentInterface) intraprocessService {
	return intraprocessService{intraprocessInterface}
}

func (i intraprocessService) StudentByID(ctx context.Context, id uint64) (domain.Student, error) {
	domainStudent, err := i.intraprocessInterface.StudentByID(ctx, id)
	if err != nil {
		return domain.Student{}, err
	}

	return domain.Student{
		ID:           domainStudent.ID,
		JenjangID:    domainStudent.JenjangID,
		ProdiID:      domainStudent.ProdiID,
		JenjangNama:  domainStudent.JenjangNama,
		ProdiNama:    domainStudent.ProdiNama,
		NPM:          domainStudent.NPM,
		Email:        domainStudent.Email,
		Nama:         domainStudent.Nama,
		Kelamin:      domainStudent.Kelamin,
		TanggalLahir: domainStudent.TanggalLahir,
	}, err
}
