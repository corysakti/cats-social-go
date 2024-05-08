CREATE TABLE cat
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    race VARCHAR(200) NOT NULL,
    sex VARCHAR(200) NOT NULL,
    description varchar(200) not null,
    age_in_month VARCHAR(200) NOT NULL,
    image_Urls VARCHAR(5000) NOT NULL,
    created_At TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);