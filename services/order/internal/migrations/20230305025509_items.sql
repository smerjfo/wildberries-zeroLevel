-- +goose Up
CREATE TABLE IF NOT EXISTS items(
    chrt_id int NOT NULL PRIMARY KEY,
    track_number varchar(255) NOT NULL,
    price int NOT NULL,
    rid varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    sale int NOT NULL,
    size varchar(255) NOT NULL,
    total_price int NOT NULL,
    nm_id int NOT NULL,
    brand varchar(255) NOT NULL,
    status int NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
