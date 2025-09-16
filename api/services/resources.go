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
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_RESOURCES", Err: err }
	}
	return resources, nil
}

func GetResourceById(db *sql.DB, id int64) (database.Resource, error) {
	row := db.QueryRow(database.RESOURCE_BY_ID, id)
	var resource database.Resource
	rowErr := row.Scan(
		&resource.ID, &resource.Link, &resource.TypeID, &resource.TopicID,
		&resource.IsValid, &resource.CreatedAt, &resource.UpdatedAt)
	if rowErr != nil {
		if rowErr == sql.ErrNoRows {
			return database.Resource{}, NotFoundError{
				Query: "RESOURCE_BY_ID", Field: "id", Value: id }
		}
		return database.Resource{}, FillColumnsError{ Err: rowErr }
	}
	return resource, nil
}

func GetAllResourcesWithDetails(db *sql.DB) ([]database.ResourceWithDetails, error) {
	rows, rowsErr := db.Query(database.ALL_RESOURCES_WITH_DETAILS)
	if rowsErr != nil {
		return nil, BadQueryError{ Query: "ALL_RESOURCES_WITH_DETAILS", Err: rowsErr }
	}
	defer rows.Close()
	var resourcesWithDetails []database.ResourceWithDetails
	for rows.Next() {
		var resourceWithDetails database.ResourceWithDetails
		rowErr := rows.Scan(
			&resourceWithDetails.ResourceID, &resourceWithDetails.ResourceLink,
			&resourceWithDetails.TypeName, &resourceWithDetails.TopicName)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		resourcesWithDetails = append(resourcesWithDetails, resourceWithDetails)
	}
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_RESOURCES_WITH_DETAILS", Err: err }
	}
	return resourcesWithDetails, nil
}

func GetResourceWithDetailsById(db *sql.DB, id int64) (database.ResourceWithDetails, error) {
	row := db.QueryRow(database.RESOURCE_WITH_DETAILS_BY_ID, id)
	var resourceWithDetails database.ResourceWithDetails
	rowErr := row.Scan(
		&resourceWithDetails.ResourceID, &resourceWithDetails.ResourceLink,
		&resourceWithDetails.TypeName, &resourceWithDetails.TopicName)
	if rowErr != nil {
		if rowErr == sql.ErrNoRows {
			return database.ResourceWithDetails{}, NotFoundError{
				Query: "RESOURCE_WITH_DETAILS_BY_ID", Field: "id", Value: id }
		}
		return database.ResourceWithDetails{}, FillColumnsError{ Err: rowErr }
	}
	return resourceWithDetails, nil
}

func GetResourcesWithDetailsByTopic(
	db *sql.DB, topicID int64) ([]database.ResourceWithDetails, error) {
	rows, rowsErr := db.Query(database.RESOURCES_BY_TOPIC, topicID)
	if rowsErr != nil { return nil, BadQueryError{ Query: "RESOURCES_BY_TOPIC", Err: rowsErr } }
	defer rows.Close()
	var resourcesByTopic []database.ResourceWithDetails
	for rows.Next() {
		var resourceByTopic database.ResourceWithDetails
		rowErr := rows.Scan(
			&resourceByTopic.ResourceID, &resourceByTopic.ResourceLink,
			&resourceByTopic.TypeName, &resourceByTopic.TopicName)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		resourcesByTopic = append(resourcesByTopic, resourceByTopic)
	}
	err := rows.Err()
	if err != nil { return nil, RowsIterationError{ Query: "RESOURCES_BY_TOPIC", Err: err } }
	if len(resourcesByTopic) == 0 {
		return nil, NotFoundError{
			Query: "RESOURCES_BY_TOPIC", Field: "topic_id", Value: topicID }
	}
	return resourcesByTopic, nil
}
