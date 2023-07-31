CREATE DATABASE wb;
CREATE USER supet1337 WITH SUPERUSER CREATEDB PASSWORD 'Asdfgh12345';
GRANT ALL PRIVILEGES ON DATABASE wb TO supet1337;
\c wb
CREATE TABLE IF NOT EXISTS orders
(
    order_uid text primary key,
    model json  not null
);