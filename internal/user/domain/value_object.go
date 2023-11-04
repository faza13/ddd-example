package domain

import "errors"

var (
	student = UserType{2}
	lecture = UserType{4}
)

var userTypeValues = []UserType{student, lecture}

func NewUserTypeFromInt(ut int8) (UserType, error) {
	for _, typeValue := range userTypeValues {
		if typeValue.Integer() == ut {
			return typeValue, nil
		}
	}

	return UserType{}, errors.New("user bukan mahasiswa/dosen")
}

type UserType struct {
	ut int8
}

func (h UserType) IsZero() bool {
	return h == UserType{}
}

func (h UserType) Integer() int8 {
	return h.ut
}
