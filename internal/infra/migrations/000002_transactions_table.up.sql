DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'transactions')THEN
    CREATE TABLE transactions(
  transaction_id SERIAL PRIMARY KEY,
  stock_id INTEGER REFERENCES stock(stock_id) NOT NULL,
  transaction_type VARCHAR(10) NOT NULL CHECK (transaction_type IN ('DEPOSIT', 'BUY', 'SELL')),
  total_amount DECIMAL(15, 2) NOT NULL,
  quantity INTEGER,
  transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
END IF;
END $$;