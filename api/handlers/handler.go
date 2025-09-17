package handlers

import "database/sql"

type Handler struct {
	db *sql.DB
}

func New(database *sql.DB) *Handler {
	return &Handler{ db: database }
}
