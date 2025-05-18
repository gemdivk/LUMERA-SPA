CREATE TABLE schedule_templates (
  id UUID PRIMARY KEY,
  specialist_id UUID NOT NULL,
  weekday INT NOT NULL,
  start_time VARCHAR(10) NOT NULL,
  end_time VARCHAR(10) NOT NULL,
  break_minutes INT DEFAULT 15
);
