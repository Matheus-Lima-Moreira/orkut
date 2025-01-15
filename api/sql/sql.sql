CREATE DATABASE IF NOT EXISTS orkut;
USE orkut;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id int auto_increment primary key,
  name varchar(50) NOT NULL,
  nick varchar(50) NOT NULL unique,
  email varchar(50) NOT NULL unique,
  password varchar(20) NOT NULL unique,
  created_at timestamp DEFAULT current_timestamp()
) ENGINE=InnoDB