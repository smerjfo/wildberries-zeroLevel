-- +goose Up
CREATE TABLE IF NOT EXISTS order_items
(
    order_uid VARCHAR(255) NOT NULL,
    chrt_id   int          NOT NULL,
    PRIMARY KEY (order_uid, chrt_id),
    FOREIGN KEY (order_uid) REFERENCES orders (order_uid) ON DELETE CASCADE,
    FOREIGN KEY (chrt_id) REFERENCES items (chrt_id) ON DELETE CASCADE
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
