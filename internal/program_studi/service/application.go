package service

import (
	"app/internal/program_studi/infrastructure/program_studi"
	"app/internal/program_studi/usecase"
	"app/pkg/database"
)

func NewApplication() usecase.Application {
	psRepo := program_studi.NewOrmRepository(database.Orm)
	return usecase.Application{List: usecase.NewListHandle(psRepo)}
}
