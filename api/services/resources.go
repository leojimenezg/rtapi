package services

import (
	"database/sql"
	"github.com/leojimenezg/rtapi/api/database"
)

func GetAllResources(db *sql.DB) ([]database.Resource, error) {
	rows, rowsErr := db.Query(database.ALL_RESOURCES)
	if rowsErr != nil { return nil, BadQueryError{ Query: "ALL_RESOURCES", Err: rowsErr } }
	defer rows.Close()
	var resources []database.Resource
	for rows.Next() {
		var resource database.Resource
		rowErr := rows.Scan(
			&resource.ID, &resource.Link, &resource.TypeID, &resource.TopicID,
			&resource.IsValid, &resource.CreatedAt, &resource.UpdatedAt)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		resources = append(resources, resource)
	}
	if err := rows.Err(); err != nil { return nil, RowsIterationError{ Query: "ALL_RESOURCES", Err: err } }
	return resources, nil
}

func GetResourceById(db *sql.DB, id int64) (database.Resource, error) {
	row := db.QueryRow(database.RESOURCE_BY_ID, id)
	var resource database.Resource
	rowErr := row.Scan(
		&resource.ID, &resource.Link, &resource.TypeID, &resource.TopicID,
		&resource.IsValid, &resource.CreatedAt, &resource.UpdatedAt)
	if rowErr != nil { return database.Resource{}, FillColumnsError{ Err: rowErr } }
	return resource, nil
}
