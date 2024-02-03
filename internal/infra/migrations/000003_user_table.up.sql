DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
CREATE TABLE users (
                       user_id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       wallet_balance DECIMAL(15, 2) DEFAULT 0.00
);
END IF;
END $$;