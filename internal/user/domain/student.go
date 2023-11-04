package domain

type Student struct {
	ID           uint64      `json:"ID"`
	JenjangID    uint64      `json:"JenjangID"`
	ProdiID      uint64      `json:"ProdiID"`
	JenjangNama  string      `json:"JenjangNama"`
	ProdiNama    string      `json:"ProdiNama"`
	NPM          string      `json:"NPM"`
	Email        string      `json:"Email"`
	Nama         string      `json:"Nama"`
	Kelamin      string      `json:"Kelamin"`
	TanggalLahir interface{} `json:"TanggalLahir"`
}
