CREATE TABLE IF NOT EXISTS orders
(
    order_uid text primary key,
    model json  not null
);