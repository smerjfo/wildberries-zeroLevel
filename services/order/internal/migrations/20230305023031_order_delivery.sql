-- +goose Up
CREATE TABLE IF NOT EXISTS order_deliveries(
    order_uid varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    phone varchar(255) NOT NULL,
    zip varchar(255) NOT NULL,
    city varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    region varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
