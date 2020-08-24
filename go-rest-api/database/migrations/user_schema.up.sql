CREATE TABLE users (
    id SERIAL UNIQUE NOT NULL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    age INT NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_on timestamp DEFAULT NULL
);