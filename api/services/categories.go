package services

import (
	"database/sql"
	"github.com/leojimenezg/rtapi/api/database"
)

func GetAllCategories(db *sql.DB) ([]database.Category, error) {
	rows, rowsErr := db.Query(database.ALL_CATEGORIES)
	if rowsErr != nil { return nil, BadQueryError{ Query: "ALL_CATEGORIES", Err: rowsErr } }
	defer rows.Close()
	var categories []database.Category
	for rows.Next() {
		var category database.Category
		rowErr := rows.Scan(
			&category.ID, &category.Name, &category.Description,
			&category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_CATEGORIES", Err: err }
	}
	return categories, nil
}

func GetCategoryById(db *sql.DB, id int64) (database.Category, error) {
	row := db.QueryRow(database.CATEGORY_BY_ID, id)
	var category database.Category
	rowErr := row.Scan(
		&category.ID, &category.Name, &category.Description,
		&category.IsActive, &category.CreatedAt, &category.UpdatedAt)
	if rowErr != nil { return database.Category{}, FillColumnsError{ Err: rowErr } }
	return category, nil
}

func GetCategoryByName(db *sql.DB, name string) (database.Category, error) {
	row := db.QueryRow(database.CATEGORY_BY_NAME, name)
	var category database.Category
	rowErr := row.Scan(
		&category.ID, &category.Name, &category.Description,
		&category.IsActive, &category.CreatedAt, &category.UpdatedAt)
	if rowErr != nil { return database.Category{}, FillColumnsError{ Err: rowErr } }
	return category, nil
}
