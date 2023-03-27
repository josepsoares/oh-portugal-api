package models

import "gorm.io/gorm"

type Portugal struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

type Region struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

type Island struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

type River struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

type Lagoon struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

type Mountain struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}

type UnescoWorldHeritageSite struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
