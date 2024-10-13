-- +migrate Up
CREATE TABLE base_trade_logs (
    block_timestamp TIMESTAMPTZ,
    block_number INTEGER,
    tx_index INTEGER, -- index of trade in this block
    tx_hash TEXT, 
    sender TEXT, 

    token_in_address TEXT,
    token_in_amount FLOAT(32),
    token_in_usdt_rate FLOAT(32),

    token_out_address TEXT,
    token_out_amount FLOAT(32),
    token_out_usdt_rate FLOAT(32),

    native_token_usdt_rate FLOAT(32),
    routes JSONB DEFAULT '{}'::jsonb,

    created TIMESTAMPTZ,
    CONSTRAINT base_trade_block_number UNIQUE (block_number, tx_index)
);

CREATE INDEX base_block_number_idx ON base_trade_logs(block_number);
CREATE INDEX base_sender_idx ON base_trade_logs(sender);
CREATE INDEX base_token_in_address_idx ON base_trade_logs(token_in_address);
CREATE INDEX base_token_out_address_idx ON base_trade_logs(token_out_address);


SELECT create_hypertable('base_trade_logs', by_range('block_number'));


CREATE TABLE base_transfer_logs (
    block_timestamp TIMESTAMPTZ,
    block_number INTEGER,
    tx_index INTEGER, -- index of trade in this block
    tx_hash TEXT, 
    
    from_address TEXT, 
    to_address TEXT,

    token_address TEXT,
    token_amount FLOAT(32),
    is_cex_in bool,
    exchange TEXT,
        
    created TIMESTAMPTZ,
    CONSTRAINT base_transfer_block_number UNIQUE (block_number, tx_index)
);

CREATE INDEX base_from_address_idx ON base_transfer_logs(from_address);
CREATE INDEX base_to_address_idx ON base_transfer_logs(to_address);
CREATE INDEX base_token_address_idx ON base_transfer_logs(token_address);

SELECT create_hypertable('base_transfer_logs', by_range('block_number'));


-- +migrate Down
DROP TABLE IF EXISTS base_trade_logs;
DROP TABLE IF EXISTS base_transfer_logs;

