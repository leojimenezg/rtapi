package services

import (
	"database/sql"
	"github.com/leojimenezg/rtapi/api/database"
)

func GetAllDifficulties(db *sql.DB) ([]database.Difficulty, error) {
	rows, rowsErr := db.Query(database.ALL_DIFFICULTIES)
	if rowsErr != nil { return nil, BadQueryError{ Query: "ALL_DIFFICULTIES", Err: rowsErr } }
	defer rows.Close()
	var difficulties []database.Difficulty
	for rows.Next() {
		var difficulty database.Difficulty
		rowErr := rows.Scan(
			&difficulty.ID, &difficulty.Name, &difficulty.Description,
			&difficulty.IsActive, &difficulty.CreatedAt, &difficulty.UpdatedAt)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		difficulties = append(difficulties, difficulty)
	}
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_DIFFICULTIES", Err: err }
	}
	return difficulties, nil
}

func GetDifficultyById(db *sql.DB, id int64) (database.Difficulty, error) {
	row := db.QueryRow(database.DIFFICULTY_BY_ID, id)
	var difficulty database.Difficulty
	rowErr := row.Scan(
		&difficulty.ID, &difficulty.Name, &difficulty.Description,
		&difficulty.IsActive, &difficulty.CreatedAt, &difficulty.UpdatedAt)
	if rowErr != nil {
		if rowErr == sql.ErrNoRows {
			return database.Difficulty{}, NotFoundError{
				Query: "DIFFICULTY_BY_ID", Field: "id", Value: id }
		}
		return database.Difficulty{}, FillColumnsError{ Err: rowErr }
	}
	return difficulty, nil
}

func GetDifficultyByName(db *sql.DB, name string) (database.Difficulty, error) {
	row := db.QueryRow(database.DIFFICULTY_BY_NAME, name)
	var difficulty database.Difficulty
	rowErr := row.Scan(
		&difficulty.ID, &difficulty.Name, &difficulty.Description,
		&difficulty.IsActive, &difficulty.CreatedAt, &difficulty.UpdatedAt)
	if rowErr != nil {
		if rowErr == sql.ErrNoRows {
			return database.Difficulty{}, NotFoundError{
				Query: "DIFFICULTY_BY_NAME", Field: "name", Value: name }
		}
		return database.Difficulty{}, FillColumnsError{ Err: rowErr }
	}
	return difficulty, nil
}
