-- +goose Up
CREATE TABLE passwords (
    id BYTEA(16) NOT NULL,
    owner_id BYTEA(16) NOT NULL,
    pass_hash CHAR(64) NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT 'false',
    PRIMARY KEY (id),
    CONSTRAINT fk_passwords_owner_id
        FOREIGN KEY (owner_id) REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE passwords;
