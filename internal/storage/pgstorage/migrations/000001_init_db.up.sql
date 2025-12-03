CREATE TABLE IF NOT EXISTS Users
(
    id                 SERIAL PRIMARY KEY,
    passwd_hashed_salt VARCHAR(1023) NOT NULL,
    tigr_id            VARCHAR(30)   NOT NULL,
    rating             INT
)