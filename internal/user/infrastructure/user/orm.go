package user

import (
	"app/internal/user/domain"
	"context"
	"gorm.io/gorm"
)

const tableName = "user"

type OrmRepository struct {
	db *gorm.DB
}

func NewOrmRepository(db *gorm.DB) OrmRepository {
	return OrmRepository{
		db: db,
	}
}

func (orm OrmRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	user := domain.User{}
	res := orm.db.WithContext(ctx).Table(tableName).Where("Nama = ?", username).First(&user)

	return user, res.Error
}
