-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE currency_rates(
    currency char(3) NOT NULL,
    timestamp timestamp NOT NULL,
    rate decimal NOT NULL,
    CONSTRAINT currency_rates_pk PRIMARY KEY (currency, timestamp)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists currency_rates;