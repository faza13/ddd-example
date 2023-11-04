package domain

type Lecture struct {
	ID           uint64      `json:"id"`
	Title        string      `json:"Title"`
	Gelar        string      `json:"Gelar"`
	Nama         string      `json:"Nama"`
	Kelamin      string      `json:"Kelamin"`
	Telepon      string      `json:"Telepon"`
	HP           string      `json:"HP"`
	Email        string      `json:"Email"`
	TanggalLahir interface{} `json:"TanggalLahir"`
}
