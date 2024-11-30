package model

type App struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}
