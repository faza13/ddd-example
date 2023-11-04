package program_studi

import (
	"app/internal/program_studi/domain"
	"context"
	"gorm.io/gorm"
)

const tableName = "ProgramStudi"

type OrmRepository struct {
	db *gorm.DB
}

func NewOrmRepository(db *gorm.DB) OrmRepository {
	return OrmRepository{
		db: db,
	}
}

func (orm OrmRepository) List(ctx context.Context) ([]domain.ProgramStudi, error) {
	data := []domain.ProgramStudi{}
	res := orm.db.WithContext(ctx).Table(tableName).Find(&data)

	return data, res.Error
}
