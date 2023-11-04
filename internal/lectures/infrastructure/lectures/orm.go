package lectures

import (
	"app/internal/lectures/domain"
	"context"
	"gorm.io/gorm"
)

const tableName = "dosen"

type OrmRepository struct {
	db *gorm.DB
}

func NewOrmRepository(db *gorm.DB) OrmRepository {
	return OrmRepository{
		db: db,
	}
}

func (orm OrmRepository) GetByID(ctx context.Context, id uint64) (domain.Lecture, error) {
	//user := domain.Lec{}
	dosen := domain.Lecture{}
	res := orm.db.WithContext(ctx).Table(tableName).Where("id = ?", id).First(&dosen)

	return dosen, res.Error
}
