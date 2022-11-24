-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE user_spendings(
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    category varchar(100) NOT NULL,
    amount decimal NOT NULL,
    date date NOT NULL
);

CREATE INDEX IF NOT EXISTS user_spendings_user_date_idx ON user_spendings(user_id, date);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table if exists user_spendings;