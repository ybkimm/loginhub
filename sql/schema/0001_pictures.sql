-- +goose Up
CREATE TABLE pictures (
    id CHAR(64) PRIMARY KEY,
    desc_text VARCHAR(160) NOT NULL,
    file_size INT NOT NULL (file_size > 0),
    width INT NOT NULL CHECK (width > 0),
    height INT NOT NULL CHECK (height > 0)
);

-- +goose Down
DROP TABLE pictures;
