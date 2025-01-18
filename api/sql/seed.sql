INSERT INTO users (name, nick, email, password)
VALUES 
("John", "John", "John@gmail.com", "$2a$10$.REUB6YbWBAaHKXo0fKYAucJK9jbRX027/0ILBttgGYE6TCpp7Fla"),
("Tom", "Tom", "Tom@gmail.com", "$2a$10$.REUB6YbWBAaHKXo0fKYAucJK9jbRX027/0ILBttgGYE6TCpp7Fla"),
("Alex", "Alex", "Alex@gmail.com", "$2a$10$.REUB6YbWBAaHKXo0fKYAucJK9jbRX027/0ILBttgGYE6TCpp7Fla");
-- all passwords = 123456

INSERT INTO followers (user_id, follower_id)
VALUES 
(1, 2),
(3, 1),
(1, 3);

INSERT INTO posts (title, content, author_id)
VALUES
("My First Post", "Hello everyone! This is my first post here. Excited to connect with all of you!", 1),
("Tech Trends 2025", "AI, blockchain, and green technology are shaping the future. What are your thoughts?", 2),
("Life Update", "Just moved to a new city and started a new job. Feeling both excited and nervous!", 3);