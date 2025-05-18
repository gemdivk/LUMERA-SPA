CREATE TABLE bookings (
  id UUID PRIMARY KEY,
  client_id UUID NOT NULL,
  specialist_id UUID NOT NULL,
  procedure_id UUID NOT NULL,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at TIMESTAMP NOT NULL
);
