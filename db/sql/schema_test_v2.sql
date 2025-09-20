-- Random Topic API Testing Database Schema
CREATE DATABASE IF NOT EXISTS rtapi_test_db;

USE rtapi_test_db;

-- Categories table
CREATE TABLE categories (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_categories_active (is_active),
    INDEX idx_categories_name (name)
);

-- Difficulties table
CREATE TABLE difficulties (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_difficulties_active (is_active),
    INDEX idx_difficulties_name (name)
);

-- Resource types table
CREATE TABLE resource_types (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_resource_types_name (name)
);

-- Topics table
CREATE TABLE topics (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL UNIQUE,
    description TEXT,
    count INT NOT NULL DEFAULT 0,
    category_id INT NOT NULL,
    difficulty_id INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (difficulty_id) REFERENCES difficulties(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    INDEX idx_topics_active (is_active),
    INDEX idx_topics_category (category_id),
    INDEX idx_topics_difficulty (difficulty_id),
    INDEX idx_topics_category_difficulty (category_id, difficulty_id),
    FULLTEXT INDEX idx_topics_search (name, description)
);

-- Resources table
CREATE TABLE resources (
    id INT NOT NULL AUTO_INCREMENT,
    link VARCHAR(700) NOT NULL UNIQUE,
    type_id INT NOT NULL,
    topic_id INT NOT NULL,
    is_valid BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (type_id) REFERENCES resource_types(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE ON UPDATE CASCADE,
    INDEX idx_resources_topic (topic_id),
    INDEX idx_resources_type (type_id),
    INDEX idx_resources_valid (is_valid)
);
