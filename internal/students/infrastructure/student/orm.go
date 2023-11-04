package student

import (
	"app/internal/students/domain"
	"context"
	"gorm.io/gorm"
)

const tableName = "mahasiswa"

type OrmRepository struct {
	db *gorm.DB
}

func NewOrmRepository(db *gorm.DB) OrmRepository {
	return OrmRepository{
		db: db,
	}
}

func (orm OrmRepository) ByID(ctx context.Context, id uint64) (domain.Student, error) {
	//user := domain.Lec{}
	student := domain.Student{}
	res := orm.db.WithContext(ctx).Table(tableName).
		Joins("left join programstudi on programstudi.ID = mahasiswa.ProdiID").
		Joins("left join jenjang on jenjang.ID = mahasiswa.JenjangID").
		Select("mahasiswa.*", "programstudi.Nama ProdiNama", "jenjang.Nama JenjangNama").
		Where("mahasiswa.ID = ?", id).First(&student)

	return student, res.Error
}
