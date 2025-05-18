CREATE TABLE daily_schedules (
  id UUID PRIMARY KEY,
  specialist_id UUID NOT NULL,
  date DATE NOT NULL,
  start_time VARCHAR(10),
  end_time VARCHAR(10),
  override BOOLEAN DEFAULT FALSE,
  cancelled BOOLEAN DEFAULT FALSE
);
