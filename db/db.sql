CREATE TABLE IF NOT EXISTS users (
    id STRING NOT NULL,
    profile_url STRING,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS links (
    userid STRING,
    longurl STRING NOT NULL,
    shorturl STRING NOT NULL,
    CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES users
)

ALTER TABLE users
RENAME COLUMN "profile url" TO profile_url;