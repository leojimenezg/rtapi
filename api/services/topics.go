package services

import (
	"database/sql"
	"github.com/leojimenezg/rtapi/api/database"
)

func GetAllTopics(db *sql.DB) ([]database.Topic, error) {
	rows, rowsErr := db.Query(database.ALL_TOPICS)
	if rowsErr != nil { return nil, BadQueryError{ Query: "ALL_TOPICS", Err: rowsErr } }
	defer rows.Close()
	var topics []database.Topic
	for rows.Next() {
		var topic database.Topic
		rowErr := rows.Scan(
			&topic.ID, &topic.Name, &topic.Description, &topic.Count, &topic.CategoryID,
			&topic.DifficultyID, &topic.IsActive, &topic.CreatedAt, &topic.UpdatedAt)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		topics = append(topics, topic)
	}
	if err := rows.Err(); err != nil { return nil, RowsIterationError{ Query: "ALL_TOPICS", Err: err } }
	return topics, nil
}

func GetTopicById(db *sql.DB, id int64) (database.Topic, error) {
	row := db.QueryRow(database.TOPIC_BY_ID, id)
	var topic database.Topic
	rowErr := row.Scan(
		&topic.ID, &topic.Name, &topic.Description, &topic.Count, &topic.CategoryID,
		&topic.DifficultyID, &topic.IsActive, &topic.CreatedAt, &topic.UpdatedAt)
	if rowErr != nil { return database.Topic{}, FillColumnsError{ Err: rowErr } }
	return topic, nil
}
