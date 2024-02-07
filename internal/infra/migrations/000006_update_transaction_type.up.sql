ALTER TABLE transactions
ADD CONSTRAINT transaction_type_check CHECK (transaction_type IN ('DEPOSIT', 'BUY', 'SELL', 'WITHDRAW'));
