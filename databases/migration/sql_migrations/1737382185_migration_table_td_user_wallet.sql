-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE tr_users_topups(
    top_up_id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    amount          DECIMAL(10,2),
    balance_before  DECIMAL(10,2),
    balance_after   DECIMAL(10,2),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(100) DEFAULT NULL,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      VARCHAR(100) DEFAULT NULL
);

CREATE TABLE tr_users_transfers(
    transfer_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID,
    user_id_target  UUID,
    status          VARCHAR(100),
    amount          DECIMAL(10,2),
    remarks         VARCHAR(100),
    balance_before  DECIMAL(10,2),
    balance_after   DECIMAL(10,2),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by      VARCHAR(100) DEFAULT NULL,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      VARCHAR(100) DEFAULT NULL
);

-- +migrate StatementEnd