SELECT * FROM categories;

SELECT * FROM difficulties;

SELECT * FROM resource_types;

SELECT * FROM topics;

SELECT * FROM resources;

-- Random topic without filters
SELECT 
    t.id AS topic_id,
    t.name AS topic_name,
    t.description AS topic_description,
    c.name AS category_name,
    c.description AS category_description,
    d.name AS difficulty_name,
    d.description AS difficulty_description
FROM topics t
JOIN categories c ON c.id = t.category_id
JOIN difficulties d ON d.id = t.difficulty_id
WHERE t.is_active = TRUE
ORDER BY RAND()
LIMIT 1;

-- Random topic with category filter
SELECT 
    t.id AS topic_id,
    t.name AS topic_name,
    t.description AS topic_description,
    c.name AS category_name,
    c.description AS category_description,
    d.name AS difficulty_name,
    d.description AS difficulty_description
FROM topics t
JOIN categories c ON c.id = t.category_id
JOIN difficulties d ON d.id = t.difficulty_id
WHERE t.is_active = TRUE
AND c.name = "Technology"
ORDER BY RAND()
LIMIT 1;

-- Random topic with difficulty filter
SELECT 
    t.id AS topic_id,
    t.name AS topic_name,
    t.description AS topic_description,
    c.name AS category_name,
    c.description AS category_description,
    d.name AS difficulty_name,
    d.description AS difficulty_description
FROM topics t
JOIN categories c ON c.id = t.category_id
JOIN difficulties d ON d.id = t.difficulty_id
WHERE t.is_active = TRUE
AND d.name = "Beginner"
ORDER BY RAND()
LIMIT 1;

-- Random topic with category and difficulty filters
SELECT 
    t.id AS topic_id,
    t.name AS topic_name,
    t.description AS topic_description,
    c.name AS category_name,
    c.description AS category_description,
    d.name AS difficulty_name,
    d.description AS difficulty_description
FROM topics t
JOIN categories c ON c.id = t.category_id
JOIN difficulties d ON d.id = t.difficulty_id
WHERE t.is_active = TRUE
AND c.name = "Technology"
AND d.name = "Beginner"
ORDER BY RAND()
LIMIT 1;

