CREATE DATABASE IF NOT EXISTS orkut;
USE orkut;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id int auto_increment primary key,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL unique,
  email varchar(50) NOT NULL unique,
  password text NOT NULL,
  created_at timestamp DEFAULT current_timestamp()
) ENGINE=InnoDB;

CREATE TABLE followers (
  user_id int not null,  
  follower_id int not null,

  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  FOREIGN KEY (follower_id)
  REFERENCES users(id)
  ON DELETE CASCADE,
  
  PRIMARY KEY (user_id, follower_id)
) ENGINE=InnoDB;

CREATE TABLE posts (
  id int auto_increment primary key,
  title text NOT NULL,
  content text NOT NULL,
  author_id int not null,
  likes int default 0,
  created_at timestamp DEFAULT current_timestamp(),

  FOREIGN KEY (author_id)
  REFERENCES users(id)
  ON DELETE CASCADE
) ENGINE=InnoDB;