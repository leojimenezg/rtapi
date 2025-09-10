-- Random Topic API Sample Data
USE rtapi_db;

-- Insert Categories
INSERT INTO categories (name, description) VALUES
("Technology", "Topics related to programming, software development, and digital technologies"),
("Science", "Topics covering natural sciences, physics, chemistry, biology, and mathematics"),
("Arts", "Topics about visual arts, music, literature, and creative expression"),
("History", "Topics covering historical events, civilizations, and cultural heritage"),
("Business", "Topics related to entrepreneurship, economics, marketing, and business strategy");

-- Insert Difficulties
INSERT INTO difficulties (name, description) VALUES
("Beginner", "No previous knowledge required. Perfect for starting a new subject"),
("Intermediate", "Basic understanding recommended. Suitable for expanding existing knowledge"),
("Advanced", "Strong foundation required. For deepening expertise in complex topics"),
("Expert", "Extensive background needed. Highly specialized and technical content");

-- Insert Resource Types
INSERT INTO resource_types (type, description) VALUES
("Article", "Written articles, blog posts, and documentation"),
("Video", "Educational videos, tutorials, and video courses"),
("Book", "Books, eBooks, and comprehensive written materials"),
("Course", "Structured online courses and learning programs"),
("Tutorial", "Step-by-step guides and hands-on tutorials"),
("Podcast", "Audio content and educational podcasts");

-- Insert Topics
INSERT INTO topics (name, description, category_id, difficulty_id) VALUES
-- Technology Topics
("Introduction to Python Programming", "Learn the basics of Python programming language including syntax, data types, and control structures", 1, 1),
("Machine Learning Fundamentals", "Understanding core concepts of ML including supervised, unsupervised learning and neural networks", 1, 2),
("Microservices Architecture", "Design patterns and best practices for building scalable microservices systems", 1, 3),
("Quantum Computing Algorithms", "Advanced algorithms and mathematical foundations of quantum computing", 1, 4),
-- Science Topics
("Solar System Basics", "Introduction to planets, moons, and celestial bodies in our solar system", 2, 1),
("Organic Chemistry Reactions", "Common reaction mechanisms and synthetic pathways in organic chemistry", 2, 2),
("Quantum Mechanics Principles", "Mathematical formalism and physical interpretations of quantum mechanics", 2, 3),
("Advanced String Theory", "Mathematical framework and physical implications of string theory in particle physics", 2, 4),
-- Arts Topics
("Color Theory in Design", "Understanding color relationships, harmony, and psychological effects in visual design", 3, 1),
("Renaissance Art History", "Major artists, techniques, and cultural context of Renaissance period art", 3, 2),
("Advanced Music Composition", "Complex harmonic structures, counterpoint, and contemporary composition techniques", 3, 3),
("Postmodern Literary Criticism", "Theoretical frameworks and analytical approaches in postmodern literature", 3, 4),
-- History Topics
("World War II Overview", "Major events, causes, and consequences of the Second World War", 4, 1),
("Medieval European Society", "Social structures, economics, and daily life in medieval Europe", 4, 2),
("Ancient Roman Politics", "Complex political systems, governance, and power structures in ancient Rome", 4, 3),
("Byzantine Historiography", "Critical analysis of historical sources and scholarly interpretations of Byzantine Empire", 4, 4),
-- Business Topics
("Basic Entrepreneurship", "Fundamentals of starting a business including planning, funding, and market research", 5, 1),
("Digital Marketing Strategy", "Comprehensive approaches to online marketing including SEO, social media, and analytics", 5, 2),
("Financial Derivatives Trading", "Advanced trading strategies and risk management in financial derivatives markets", 5, 3),
("Behavioral Economics Theory", "Psychological factors in economic decision-making and market behavior analysis", 5, 4);

-- Insert Resources
INSERT INTO resources (link, type_id, topic_id) VALUES
-- Python Programming Resources
("https://docs.python.org/3/tutorial/", 1, 1),
("https://www.youtube.com/watch?v=eWRfhZUzrAc", 2, 1),
("https://automatetheboringstuff.com/", 3, 1),
-- Machine Learning Resources
("https://scikit-learn.org/stable/tutorial/index.html", 1, 2),
("https://www.coursera.org/learn/machine-learning", 4, 2),
("https://www.youtube.com/playlist?list=PLblh5JKOoLUIzaEkCLIUxQFjPIlapw8nU", 2, 2),
-- Solar System Resources
("https://solarsystem.nasa.gov/", 1, 5),
("https://www.khanacademy.org/science/cosmology-and-astronomy", 4, 5),
("https://www.youtube.com/watch?v=libKVRa01L8", 2, 5),
-- Color Theory Resources
("https://www.interaction-design.org/literature/topics/color-theory", 1, 9),
("https://www.youtube.com/watch?v=Qj1FK8n7WgY", 2, 9),
("https://www.amazon.com/Interaction-Color-Anniversary-Josef-Albers/dp/0300179359", 3, 9),
-- Entrepreneurship Resources
("https://www.sba.gov/business-guide", 1, 17),
("https://www.ted.com/playlists/163/the_entrepreneurial_mind", 2, 17),
("https://www.coursera.org/learn/entrepreneurship", 4, 17);

