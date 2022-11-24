-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE user_spend_limit_settings(
    user_id int NOT NULL,
    period_month int NOT NULL,
    spend_limit decimal NOT NULL,
    CONSTRAINT user_spend_limit_settings_pk PRIMARY KEY (user_id, period_month)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists user_spend_limit_settings;