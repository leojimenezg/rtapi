SET SQL_SAFE_UPDATES = 0;

-- Random Topic API Testing Database Seed Data
USE rtapi_test_db;

-- Clear existing data
DELETE FROM resources;
DELETE FROM topics;
DELETE FROM categories;
DELETE FROM difficulties;
DELETE FROM resource_types;

-- Reset AUTO_INCREMENT
ALTER TABLE resources AUTO_INCREMENT = 1;
ALTER TABLE topics AUTO_INCREMENT = 1;
ALTER TABLE categories AUTO_INCREMENT = 1;
ALTER TABLE difficulties AUTO_INCREMENT = 1;
ALTER TABLE resource_types AUTO_INCREMENT = 1;

-- Insert Test Categories
INSERT INTO categories (name, description) VALUES
("Technology", "Test category for tech topics"),
("Science", "Test category for science topics"),
("TestCategory", "Special category for testing edge cases");

-- Insert Test Difficulties  
INSERT INTO difficulties (name, description) VALUES
("Beginner", "Test difficulty - beginner level"),
("Advanced", "Test difficulty - advanced level"),
("TestDifficulty", "Special difficulty for testing");

-- Insert Test Resource Types
INSERT INTO resource_types (name, description) VALUES
("Article", "Test resource type - articles"),
("Video", "Test resource type - videos"),
("TestType", "Special resource type for testing");

-- Insert Test Topics
INSERT INTO topics (name, description, category_id, difficulty_id) VALUES
("Test Python Basics", "Testing topic for Python programming", 1, 1),
("Test Data Structures", "Testing topic for data structures and algorithms", 1, 2),
("Test Physics Concepts", "Testing topic for basic physics", 2, 1),
("Test Topic with Special Characters !@#", "Topic with special characters for edge case testing", 3, 3),
("Test Very Long Topic Name That Exceeds Normal Length But Is Still Valid According To Schema", "Long topic name for boundary testing", 1, 1);

-- Insert Test Resources
INSERT INTO resources (link, type_id, topic_id) VALUES
("https://test-python-article.com", 1, 1),
("https://test-python-video.com", 2, 1),
("https://test-datastructures-article.com", 1, 2),
("https://test-physics-video.com", 2, 3),
("https://test-special-chars-resource.com/?param=value&test=123", 3, 4);

-- Update topic counts based on resources
UPDATE topics SET count = (
    SELECT COUNT(*) FROM resources WHERE topic_id = topics.id
);
