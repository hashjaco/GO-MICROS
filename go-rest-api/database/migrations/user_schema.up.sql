CREATE DATABASE users_db;
USE users_db;

CREATE TABLE users (
    id INT UNIQUE NOT NULL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    username VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(30) NOT NULL,
    created_on VARCHAR(30) NOT NULL,
    updated_on VARCHAR(30) NOT NULL,
);

LOAD DATA LOCAL INFILE '../data/users.txt' INTO TABLE users;