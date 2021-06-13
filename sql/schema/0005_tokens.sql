-- +goose Up
CREATE TABLE tokens (
    id CHAR(64) NOT NULL,
    owner_id CHAR(64) NOT NULL,
    client_id CHAR(64) NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT 'false',
    device_name VARCHAR(128) NOT NULL,
    country_name VARCHAR(128) NOT NULL,
    last_access TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tokens_owner_id
        FOREIGN KEY (owner_id) REFERENCES users(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_tokens_client_id
        FOREIGN KEY (client_id) REFERENCES clients(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE tokens;
