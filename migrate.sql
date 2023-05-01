-- +migrate Up

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(36) PRIMARY KEY,
    owned_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    wallet_id VARCHAR(36) NOT NULL,
    amount BIGINT NOT NULL,
    create_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS wallets (
    id VARCHAR(36) PRIMARY KEY,
    owned_by VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    enabled_at TIMESTAMP WITH TIME ZONE NOT NULL,
    balance BIGINT NOT NULL,
    create_at TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(36) PRIMARY KEY,
    customer_xid VARCHAR(36) NOT NULL,
    token VARCHAR(255) NOT NULL,
    create_at TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS accounts;