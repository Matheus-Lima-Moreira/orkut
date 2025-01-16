CREATE DATABASE IF NOT EXISTS orkut;
USE orkut;

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