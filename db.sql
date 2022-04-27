CREATE TABLE users (
    id STRING NOT NULL,
    PRIMARY KEY (id),
);

CREATE TABLE links (
    id STRING NOT NULL,
    original_link STRING NOT NULL,
    shortened_link STRING NOT NULL,
    user_id STRING NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_userid FOREIGN KEY (user_id) REFERENCES users
);