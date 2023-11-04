package usecase

import (
	"app/internal/user/domain"
	"context"
	"errors"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type lectureServiceInterface interface {
	LectureByID(ctx context.Context, id uint64) (domain.Lecture, error)
}

type studentServiceInterface interface {
	StudentByID(ctx context.Context, id uint64) (domain.Student, error)
}

type loginHandle struct {
	userRepository domain.UserRepository
	lectureService lectureServiceInterface
	studentService studentServiceInterface
}

type LoginHandle interface {
	Handle(ctx context.Context, login LoginDTO) (any, error)
}

func NewLoginHandle(repository domain.UserRepository, lectureService lectureServiceInterface, studentService studentServiceInterface) LoginHandle {
	return loginHandle{userRepository: repository, lectureService: lectureService, studentService: studentService}
}

func (l loginHandle) Handle(ctx context.Context, login LoginDTO) (any, error) {
	user, err := l.userRepository.GetByUsername(ctx, login.Username)
	if err != nil {
		return nil, err
	}

	if !user.VerifyPassword(user.Password) {
		return nil, errors.New("username atau password salah")
	}

	if user.TabelUserID == 2 {
		return l.lectureService.LectureByID(ctx, user.EntityID)
	}

	return l.studentService.StudentByID(ctx, user.EntityID)
}
