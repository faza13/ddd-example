package intraprocess

import (
	"app/internal/students/domain"
	"context"
)

type Student struct {
	ID           uint64      `json:"ID"`
	JenjangID    uint64      `json:"JenjangID"`
	ProdiID      uint64      `json:"ProdiID"`
	JenjangNama  string      `json:"JenjangNama"`
	ProdiNama    string      `json:"ProdiNama"`
	NPM          string      `json:"NPM"`
	Email        string      `json:"Email"`
	Nama         string      `json:"Nama"`
	Kelamin      string      `json:"Kelamin"`
	TanggalLahir interface{} `json:"TanggalLahir"`
}

func StudentFromDomainStudent(domainStudent domain.Student) Student {
	return Student{
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
	}
}

type StudentInterface struct {
	repo domain.StudentRepository
}

func NewStudentInterface(repo domain.StudentRepository) StudentInterface {
	return StudentInterface{repo: repo}
}

func (i StudentInterface) StudentByID(ctx context.Context, id uint64) (Student, error) {
	domainStudent, err := i.repo.ByID(ctx, id)

	if err != nil {
		return Student{}, err
	}

	return StudentFromDomainStudent(domainStudent), err
}
