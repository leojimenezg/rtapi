package database

const (
	// -----TEMPLATE QUERIES-----
	TOPIC_WITH_DETAILS_TEMPLATE = `
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
	RESOURCE_WITH_DETAILS_TEMPLATE = `
		SELECT 
			r.id,
			r.link,
			rt.name,
			t.name
		FROM resources r
		JOIN resource_types rt ON rt.id = r.type_id
		JOIN topics t ON t.id = r.topic_id
		WHERE r.is_valid = TRUE`

	// -----QUERIES FOR CATEGORY TABLE-----
	ALL_CATEGORIES   = `SELECT * FROM categories WHERE is_active = TRUE`
	CATEGORY_BY_ID   = `SELECT * FROM categories WHERE id = ? AND is_active = TRUE`
	CATEGORY_BY_NAME = `SELECT * FROM categories WHERE name = ? AND is_active = TRUE`

	// -----QUERIES FOR DIFFICULTY TABLE-----
	ALL_DIFFICULTIES   = `SELECT * FROM difficulties WHERE is_active = TRUE`
	DIFFICULTY_BY_ID   = `SELECT * FROM difficulties WHERE id = ? AND is_active = TRUE`
	DIFFICULTY_BY_NAME = `SELECT * FROM difficulties WHERE name = ? AND is_active = TRUE`

	// -----QUERIES FOR RESOURCE_TYPES TABLE-----
	ALL_RESOURCE_TYPES    = `SELECT * FROM resource_types`
	RESOURCE_TYPE_BY_ID   = `SELECT * FROM resource_types WHERE id = ?`
	RESOURCE_TYPE_BY_NAME = `SELECT * FROM resource_types WHERE name = ?`

	// -----QUERIES FOR RESOURCES TABLE-----
	ALL_RESOURCES  = `SELECT * FROM resources WHERE is_valid = TRUE`
	RESOURCE_BY_ID = `SELECT * FROM resources WHERE id = ? AND is_valid = TRUE`

	// -----QUERIES FOR TOPICS TABLE-----
	ALL_TOPICS    = `SELECT * FROM topics WHERE is_active = TRUE`
	TOPIC_BY_ID   = `SELECT * FROM topics WHERE id = ? AND is_active = TRUE`
	TOPIC_BY_NAME = `SELECT * FROM topics WHERE name = ? AND is_active = TRUE`
	INCREMENT_TOPIC_COUNT = `
		UPDATE topics 
		SET count = count + 1 
		WHERE id = ?
		AND is_active = TRUE`

	// -----QUERIES FOR TOPICS TABLE WITH DETAILS-----
	ALL_TOPICS_WITH_DETAILS    = TOPIC_WITH_DETAILS_TEMPLATE
	TOPIC_WITH_DETAILS_BY_ID   = TOPIC_WITH_DETAILS_TEMPLATE + ` AND t.id = ?`
	TOPIC_WITH_DETAILS_BY_NAME = TOPIC_WITH_DETAILS_TEMPLATE + ` AND t.name = ?`

	// -----QUERIES FOR TOPICS TABLE WITH DETAILS RANDOMLY-----
	RANDOM_TOPIC = TOPIC_WITH_DETAILS_TEMPLATE + `
		ORDER BY RAND()
		LIMIT 1`
	RANDOM_TOPIC_BY_CATEGORY = TOPIC_WITH_DETAILS_TEMPLATE + `
		AND c.id = ?
		ORDER BY RAND()
		LIMIT 1`
	RANDOM_TOPIC_BY_DIFFICULTY = TOPIC_WITH_DETAILS_TEMPLATE + `
		AND d.id = ?
		ORDER BY RAND()
		LIMIT 1`
	RANDOM_TOPIC_BY_CATEGORY_AND_DIFFICULTY = TOPIC_WITH_DETAILS_TEMPLATE + `
		AND c.id = ?
		AND d.id = ?
		ORDER BY RAND()
		LIMIT 1`

	// -----QUERIES FOR RESOURCES TABLE WITH DETAILS-----
	ALL_RESOURCES_WITH_DETAILS    = RESOURCE_WITH_DETAILS_TEMPLATE
	RESOURCE_WITH_DETAILS_BY_ID   = RESOURCE_WITH_DETAILS_TEMPLATE + ` AND r.id = ?`
	RESOURCES_BY_TOPIC = RESOURCE_WITH_DETAILS_TEMPLATE + `AND t.id = ?`
)
