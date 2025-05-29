CREATE TABLE payments (
                          id TEXT PRIMARY KEY,
                          user_id TEXT NOT NULL,
                          amount BIGINT NOT NULL,
                          currency TEXT NOT NULL,
                          payment_method TEXT NOT NULL,
                          stripe_id TEXT NOT NULL,
                          status TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
