-- +goose Up
CREATE TABLE IF NOT EXISTS orders(
    order_uid varchar(255) NOT NULL PRIMARY KEY,
    track_number varchar(255) NOT NULL,
    entry varchar(255) NOT NULL,
    locale varchar(255) NOT NULL,
    internal_signature varchar(255) NOT NULL,
    customer_id varchar(255) NOT NULL,
    delivery_service varchar(255) NOT NULL,
    shardkey varchar(255) NOT NULL,
    sm_id int NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR(255) NOT NULL
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
