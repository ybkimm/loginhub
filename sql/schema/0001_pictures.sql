-- +goose Up
CREATE TABLE pictures (
    id BYTEA(16) NOT NULL,
    desc_text VARCHAR(160) NOT NULL,
    file_size INT NOT NULL CHECK (file_size > 0),
    width INT NOT NULL CHECK (width > 0),
    height INT NOT NULL CHECK (height > 0),
    PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE pictures;
