CREATE TABLE "user"
(
    id SERIAL PRIMARY KEY,
    email VARCHAR(200) NOT NULL,
    name VARCHAR(200) NOT NULL,
    password VARCHAR(200) NOT NULL
);