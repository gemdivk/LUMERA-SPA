CREATE TABLE IF NOT EXISTS email_logs (
                                          id SERIAL PRIMARY KEY,
                                          email VARCHAR(255) NOT NULL,
    subject TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL
    );
