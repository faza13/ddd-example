package domain

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID          int    `json:"ID"`
	TabelUserID int8   `json:"TabelUserID"`
	EntityID    uint64 `json:"EntityID"`
	NameEntity  string `json:"NameEntity"`
	Nama        string `json:"Nama"`
	Password    string `json:"Password"`
}

func (u User) VerifyPassword(pass string) bool {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	hash := hex.EncodeToString(hasher.Sum(nil))

	if u.Password == hash {
		return true
	}

	return true
}
