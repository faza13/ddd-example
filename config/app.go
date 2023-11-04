package config

import (
	"log"
	"strconv"
)

var App app

type app struct {
	Env      string
	Port     string
	Name     string
	Timezone string
}

func (a app) CheckPort() {
	_, err := strconv.ParseInt(a.Port, 10, 0)
	if err != nil {
		log.Fatal("APP Port is not numeric")
	}
}
