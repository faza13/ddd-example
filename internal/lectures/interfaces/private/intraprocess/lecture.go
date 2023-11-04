package intraprocess

import (
	"app/internal/lectures/domain"
	"context"
)

type Lecture struct {
	ID           uint64      `json:"id"`
	Title        string      `json:"Title"`
	Gelar        string      `json:"Gelar"`
	Nama         string      `json:"Nama"`
	Kelamin      string      `json:"Kelamin"`
	Telepon      string      `json:"Telepon"`
	HP           string      `json:"HP"`
	Email        string      `json:"Email"`
	TanggalLahir interface{} `json:"TanggalLahir"`
}

func LectureFromDomainLecture(domainLecture domain.Lecture) Lecture {
	return Lecture{
		ID:           domainLecture.ID,
		Title:        domainLecture.Title,
		Gelar:        domainLecture.Gelar,
		Nama:         domainLecture.Nama,
		Kelamin:      domainLecture.Kelamin,
		Telepon:      domainLecture.Telepon,
		HP:           domainLecture.HP,
		Email:        domainLecture.Email,
		TanggalLahir: domainLecture.TanggalLahir,
	}
}

type LectureInterface struct {
	repo domain.LectureRepository
}

func NewLectureInterface(repo domain.LectureRepository) LectureInterface {
	return LectureInterface{repo: repo}
}

func (i LectureInterface) LectureByID(ctx context.Context, id uint64) (Lecture, error) {
	domainLecture, err := i.repo.GetByID(ctx, id)

	if err != nil {
		return Lecture{}, err
	}

	return LectureFromDomainLecture(domainLecture), err
}
