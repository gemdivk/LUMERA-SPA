CREATE TABLE auth_users (
                            id UUID PRIMARY KEY,
                            name VARCHAR(100) UNIQUE NOT NULL,
                            email VARCHAR(255) UNIQUE NOT NULL,
                            password TEXT NOT NULL,
                            created_at TIMESTAMP NOT NULL
);
