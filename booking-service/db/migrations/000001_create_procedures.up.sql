CREATE TABLE procedures (
  id UUID PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  duration_minutes INT NOT NULL,
  created_at TIMESTAMP NOT NULL
);
