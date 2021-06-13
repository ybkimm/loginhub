-- +goose Up
CREATE TABLE clients (
    id CHAR(64) NOT NULL,
    client_name VARCHAR(128) NOT NULL,
    secret_hash CHAR(64) NOT NULL,
    owner_id CHAR(64) NOT NULL,
    client_category VARCHAR(128) NOT NULL,
    client_desc TEXT NOT NULL DEFAULT '',
    banner_picture_id CHAR(64),
    icon_id CHAR(64),
    redirect_uris TEXT,
    terms_url varchar(512),
    privacy_policy_url varchar(512),
    flag INT,
    PRIMARY KEY (id),
    CONSTRAINT fk_clients_owner_id
        FOREIGN KEY (owner_id) REFERENCES users(id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_clients_banner_picture_id
        FOREIGN KEY (banner_picture_id) REFERENCES pictures(id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_clients_icon_id
        FOREIGN KEY (icon_id) REFERENCES pictures(id)
        ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE clients
