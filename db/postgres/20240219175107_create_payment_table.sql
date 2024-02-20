-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments
(
    uuid           VARCHAR(255) PRIMARY KEY,
    status         VARCHAR(255) NOT NULL,
    description    VARCHAR(255) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    account_holder VARCHAR(255) NOT NULL,
    bank_code      VARCHAR(255) NOT NULL,
    external_id    VARCHAR(255) NOT NULL,
    reference      VARCHAR(255) NOT NULL,
    amount         INTEGER      NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at     TIMESTAMP,
    updated_at     TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
