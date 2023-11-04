package service

import (
	"app/internal/students/infrastructure/student"
	"app/internal/students/usecase"
	"app/pkg/database"
)

func NewApplication() usecase.Application {
	studentRepository := student.NewOrmRepository(database.Orm)
	return usecase.Application{
		GetByID: usecase.NewGetByIDHandle(studentRepository),
	}
}
