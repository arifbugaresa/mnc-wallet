-- +migrate Up
-- +migrate StatementBegin

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE tm_users_accounts(
    user_id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name      VARCHAR(50),
    last_name       VARCHAR(50),
    phone_number    VARCHAR(100) NOT NULL UNIQUE,
    address         VARCHAR(100),
    pin             INT,
    balance         DECIMAL(10,2),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(100) DEFAULT NULL,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      VARCHAR(100) DEFAULT NULL
);

-- +migrate StatementEnd