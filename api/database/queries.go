package database

const (
	// Basic table queries
	ALL_CATEGORIES     = `SELECT * FROM categories`
	ALL_DIFFICULTIES   = `SELECT * FROM difficulties`
	ALL_RESOURCE_TYPES = `SELECT * FROM resource_types`
	ALL_RESOURCES      = `SELECT * FROM resources`
	ALL_TOPICS         = `SELECT * FROM topics`

	// Categories
	CATEGORY_BY_ID   = `SELECT * FROM categories WHERE id = ?`
	CATEGORY_BY_NAME = `SELECT * FROM categories WHERE name = ?`

	// Difficulties
	DIFFICULTY_BY_ID   = `SELECT * FROM difficulties WHERE id = ?`
	DIFFICULTY_BY_NAME = `SELECT * FROM difficulties WHERE name = ?`

	// Resource Types
	RESOURCE_TYPE_BY_ID   = `SELECT * FROM resource_types WHERE id = ?`
	RESOURCE_TYPE_BY_NAME = `SELECT * FROM resource_types WHERE name = ?`

	// Resources
	RESOURCE_BY_ID = `SELECT * FROM resources WHERE id = ?`

	// Topics
	TOPIC_BY_ID   = `SELECT * FROM topics WHERE id = ?`
	TOPIC_BY_NAME = `SELECT * FROM topics WHERE name = ?`

	// Topic with details base query
	TOPIC_WITH_DETAILS = `
		SELECT 
			t.id,
			t.name,
			t.description,
			c.name,
			c.description,
			d.name,
			d.description
		FROM topics t
		JOIN categories c ON c.id = t.category_id
		JOIN difficulties d ON d.id = t.difficulty_id
		WHERE t.is_active = TRUE`

	// Topic with details
	TOPIC_WITH_DETAILS_BY_ID = TOPIC_WITH_DETAILS + ` AND t.id = ?`
	TOPIC_WITH_DETAILS_BY_NAME = TOPIC_WITH_DETAILS + ` AND t.name = ?`

	// Random topic
	RANDOM_TOPIC = TOPIC_WITH_DETAILS + `
		ORDER BY RAND()
		LIMIT 1`

	// Random topic by category
	RANDOM_TOPIC_BY_CATEGORY = TOPIC_WITH_DETAILS + `
		AND c.id = ?
		ORDER BY RAND()
		LIMIT 1`

	// Random topic by difficulty
	RANDOM_TOPIC_BY_DIFFICULTY = TOPIC_WITH_DETAILS + `
		AND d.id = ?
		ORDER BY RAND()
		LIMIT 1`

	// Random topic by category and difficulty
	RANDOM_TOPIC_BY_CATEGORY_AND_DIFFICULTY = TOPIC_WITH_DETAILS + `
		AND c.id = ?
		AND d.id = ?
		ORDER BY RAND()
		LIMIT 1`

	// Resource queries
	RESOURCES_BY_TOPIC = `
		SELECT 
			r.id,
			r.link,
			rt.name,
			t.name
		FROM resources r
		JOIN resource_types rt ON rt.id = r.type_id
		JOIN topics t ON t.id = r.topic_id
		WHERE r.topic_id = ?
		AND r.is_valid = TRUE`

	// Topic count increment
	INCREMENT_TOPIC_COUNT = `
		UPDATE topics 
		SET count = count + 1 
		WHERE id = ?`
)
