CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    phone VARCHAR(11) UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

DROP TABLE IF EXISTS users;

INSERT INTO users (first_name, last_name, phone, email)
    VALUES ('Ethan','Marchand','7808680552','ethan@k2r.ca');