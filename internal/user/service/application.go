package service

import (
	"app/internal/lectures/infrastructure/lectures"
	"app/internal/lectures/interfaces/private/intraprocess"
	"app/internal/students/infrastructure/student"
	studentintraprocess "app/internal/students/interfaces/private/intraprocess"
	lectures2 "app/internal/user/infrastructure/lectures"
	studentuserinfra "app/internal/user/infrastructure/students"
	"app/internal/user/infrastructure/user"
	"app/internal/user/usecase"
	"app/pkg/database"
)

func NewApplication() usecase.Application {
	lecturesRepo := lectures.NewOrmRepository(database.Orm)
	lectureInterface := intraprocess.NewLectureInterface(lecturesRepo)
	lecturesIntraprocess := lectures2.NewIntraprocessService(lectureInterface)

	studentRepo := student.NewOrmRepository(database.Orm)
	studentInterface := studentintraprocess.NewStudentInterface(studentRepo)
	studentIntraprocess := studentuserinfra.NewIntraprocessService(studentInterface)

	userRepository := user.NewOrmRepository(database.Orm)
	return usecase.Application{Login: usecase.NewLoginHandle(userRepository, lecturesIntraprocess, studentIntraprocess)}
}
