-- +goose Up
CREATE TABLE IF NOT EXISTS order_payments(
    order_uid varchar(255) NOT NULL,
    transaction varchar(255) NOT NULL,
    request_id varchar(255) NOT NULL ,
    currency varchar(255) NOT NULL ,
    provider varchar(255) NOT NULL ,
    amount int NOT NULL ,
    payment_dt TIMESTAMP NOT NULL,
    bank varchar(255) NOT NULL,
    delivery_cost int NOT NULL,
    goods_total int NOT NULL,
    custom_fee int NOT NULL,
    CONSTRAINT FK_order_payments_order_uid FOREIGN KEY (order_uid) REFERENCES orders(order_uid)

);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
