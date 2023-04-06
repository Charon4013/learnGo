package models

type Picture struct {
	Id          int    `uri:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
