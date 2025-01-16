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