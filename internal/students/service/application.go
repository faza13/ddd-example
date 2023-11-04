package service

import (
	"app/internal/lectures/infrastructure/lectures"
	"app/internal/lectures/usecase"
	"app/pkg/database"
)

func NewApplication() usecase.Application {
	lectureRepository := lectures.NewOrmRepository(database.Orm)
	return usecase.Application{
		GetByID: usecase.NewGetByIDHandle(lectureRepository),
	}
}
