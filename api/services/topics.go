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
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_TOPICS", Err: err }
	}
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

func GetTopicByName(db *sql.DB, name string) (database.Topic, error) {
	row := db.QueryRow(database.TOPIC_BY_NAME, name)
	var topic database.Topic
	rowErr := row.Scan(
		&topic.ID, &topic.Name, &topic.Description, &topic.Count, &topic.CategoryID,
		&topic.DifficultyID, &topic.IsActive, &topic.CreatedAt, &topic.UpdatedAt)
	if rowErr != nil { return database.Topic{}, FillColumnsError{ Err: rowErr } }
	return topic, nil
}

func UpdateTopicCount(db *sql.DB, id int64) error {
	_, updateErr := db.Exec(database.INCREMENT_TOPIC_COUNT, id)
	if updateErr != nil {
		return UpdateRegisterError{ Query: "INCREMENT_TOPIC_COUNT", Err: updateErr }
	}
	return nil
}

func GetAllTopicsWithDetails(db *sql.DB) ([]database.TopicWithDetails, error) {
	rows, rowsErr := db.Query(database.ALL_TOPICS_WITH_DETAILS)
	if rowsErr != nil {
		return nil, BadQueryError{ Query: "ALL_TOPICS_WITH_DETAILS", Err: rowsErr }
	}
	defer rows.Close()
	var topicsWithDetails []database.TopicWithDetails
	for rows.Next() {
		var topicWithDetails database.TopicWithDetails
		rowErr := rows.Scan(
			&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
			&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
			&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
		if rowErr != nil { return nil, FillColumnsError{ Err: rowErr } }
		topicsWithDetails = append(topicsWithDetails, topicWithDetails)
	}
	if err := rows.Err(); err != nil {
		return nil, RowsIterationError{ Query: "ALL_TOPICS_WITH_DETAILS", Err: err }
	}
	return topicsWithDetails, nil
}

func GetTopicWithDetailsById(db *sql.DB, id int64) (database.TopicWithDetails, error) {
	row := db.QueryRow(database.TOPIC_WITH_DETAILS_BY_ID, id)
	var topicWithDetails database.TopicWithDetails
	rowErr := row.Scan(
		&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
		&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
		&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
	if rowErr != nil { return database.TopicWithDetails{}, FillColumnsError{ Err: rowErr } }
	return topicWithDetails, nil
}

func GetTopicsWithDetailsByName(db *sql.DB, name string) (database.TopicWithDetails, error) {
	row := db.QueryRow(database.TOPIC_WITH_DETAILS_BY_NAME, name)
	var topicWithDetails database.TopicWithDetails
	rowErr := row.Scan(
		&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
		&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
		&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
	if rowErr != nil { return database.TopicWithDetails{}, FillColumnsError{ Err: rowErr } }
	return topicWithDetails, nil
}

func GetRandomTopicWithDetails(db *sql.DB) (database.TopicWithDetails, error) {
	row := db.QueryRow(database.RANDOM_TOPIC)
	var topicWithDetails database.TopicWithDetails
	rowErr := row.Scan(
		&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
		&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
		&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
	if rowErr != nil { return database.TopicWithDetails{}, FillColumnsError{ Err: rowErr } }
	return topicWithDetails, nil
}

func GetRandomTopicWithDetailsByCategory(
	db *sql.DB, categoryID int64) (database.TopicWithDetails, error) {
	row := db.QueryRow(database.RANDOM_TOPIC_BY_CATEGORY, categoryID)
	var topicWithDetails database.TopicWithDetails
	rowErr := row.Scan(
		&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
		&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
		&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
	if rowErr != nil { return database.TopicWithDetails{}, FillColumnsError{ Err: rowErr } }
	return topicWithDetails, nil
}

func GetRandomTopicWithDetailsByDifficulty(
	db *sql.DB, difficultyID int64) (database.TopicWithDetails, error) {
	row := db.QueryRow(database.RANDOM_TOPIC_BY_DIFFICULTY, difficultyID)
	var topicWithDetails database.TopicWithDetails
	rowErr := row.Scan(
		&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
		&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
		&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
	if rowErr != nil { return database.TopicWithDetails{}, FillColumnsError{ Err: rowErr } }
	return topicWithDetails, nil
}

func GetRandomTopicWithDetailsByCategoryAndDifficulty(
	db *sql.DB, categoryID, difficultyID int64) (database.TopicWithDetails, error) {
	row := db.QueryRow(database.RANDOM_TOPIC_BY_CATEGORY_AND_DIFFICULTY, categoryID, difficultyID)
	var topicWithDetails database.TopicWithDetails
	rowErr := row.Scan(
		&topicWithDetails.TopicID, &topicWithDetails.TopicName, &topicWithDetails.TopicDesc,
		&topicWithDetails.CategoryName, &topicWithDetails.CategoryDesc,
		&topicWithDetails.DifficultyName, &topicWithDetails.DifficultyDesc)
	if rowErr != nil { return database.TopicWithDetails{}, FillColumnsError{ Err: rowErr } }
	return topicWithDetails, nil
}
