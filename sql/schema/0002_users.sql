-- +goose Up
CREATE TABLE users (
    id CHAR(64) NOT NULL,
    email VARCHAR(512) NOT NULL,
    picture_id CHAR(64) NOT NULL,
    given_name VARCHAR(64) NOT NULL,
    family_name VARCHAR(64) NOT NULL,
    gender CHAR(1) NOT NULL,
    birthdate TIMESTAMP NOT NULL,
    flag INT NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(email),
    CONSTRAINT fk_users_picture_id
        FOREIGN KEY (picture_id) REFERENCES pictures(id)
        ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE users;
