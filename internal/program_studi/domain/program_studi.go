package domain

type ProgramStudi struct {
	ID     uint   `json:"id" gorm:"column:ID"`
	Nama   string `json:"nama" gorm:"column:Nama"`
	Status string `json:"status" gorm:"column:Status"`
}
