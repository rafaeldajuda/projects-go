package models

type Personalidade struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
