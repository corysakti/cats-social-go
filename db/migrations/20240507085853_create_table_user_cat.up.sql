CREATE TABLE user_cat_match
(
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(200) NOT NULL,
    cat_id VARCHAR(200) NOT NULL,
    receiver_cat_id VARCHAR(200) NOT NULL,
    status VARCHAR(200) NOT NULL
);
