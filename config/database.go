package config

import (
	"log"
	"strconv"
)

var Database database

type database struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (a database) CheckPort() {
	_, err := strconv.ParseInt(a.Port, 10, 0)
	if err != nil {
		log.Fatal("DB Port is not numeric")
	}
}
