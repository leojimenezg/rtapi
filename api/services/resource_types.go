package services

import (
	"database/sql"
	"github.com/leojimenezg/rtapi/api/database"
)

func GetAllResourceTypes(db *sql.DB) ([]database.ResourceType, error) {
	rows, rowsErr := db.Query(database.ALL_RESOURCE_TYPES)
	if rowsErr != nil { return nil, BadQueryError{ Query: "ALL_RESOURCE_TYPES", Err: rowsErr } }
	defer rows.Close()
	var resourceTypes []database.ResourceType
	for rows.Next() {
		var resourceType database.ResourceType
		rowErr := rows.Scan(
			&resourceType.ID, &resourceType.Name, &resourceType.Description,
			&resourceType.CreatedAt, &resourceType.UpdatedAt)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		resourceTypes = append(resourceTypes, resourceType)
	}
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_RESOURCE_TYPES", Err: err }
	}
	return resourceTypes, nil
}

func GetResourceTypeById(db *sql.DB, id int64) (database.ResourceType, error) {
	row := db.QueryRow(database.RESOURCE_TYPE_BY_ID, id)
	var resourceType database.ResourceType
	rowErr := row.Scan(
		&resourceType.ID, &resourceType.Name, &resourceType.Description,
		&resourceType.CreatedAt, &resourceType.UpdatedAt)
	if rowErr != nil { return database.ResourceType{}, FillColumnsError{ Err: rowErr } }
	return resourceType, nil
}

func GetResourceTypeByName(db *sql.DB, name string) (database.ResourceType, error) {
	row := db.QueryRow(database.RESOURCE_TYPE_BY_NAME, name)
	var resourceType database.ResourceType
	rowErr := row.Scan(
		&resourceType.ID, &resourceType.Name, &resourceType.Description,
		&resourceType.CreatedAt, &resourceType.UpdatedAt)
	if rowErr != nil { return database.ResourceType{}, FillColumnsError{ Err: rowErr } }
	return resourceType, nil
}
