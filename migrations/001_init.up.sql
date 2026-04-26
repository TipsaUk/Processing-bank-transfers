CREATE TABLE accounts (
                          id UUID PRIMARY KEY,
                          account_holder TEXT NOT NULL,
                          balance NUMERIC(18,2) NOT NULL CHECK (balance >= 0),
                          currency VARCHAR(3) NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions (
                              id UUID PRIMARY KEY,
                              from_account UUID NOT NULL,
                              to_account UUID NOT NULL,
                              amount NUMERIC(18,2) NOT NULL CHECK (amount > 0),
                              status VARCHAR(20) NOT NULL,
                              created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                              processed_at TIMESTAMP
);