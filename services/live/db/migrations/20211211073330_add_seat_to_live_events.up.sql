ALTER TABLE live_events ADD COLUMN seats INTEGER
NOT NULL CONSTRAINT positive CHECK (seats >= 0);

ALTER TABLE live_events ADD COLUMN available_seats INTEGER
NOT NULL CONSTRAINT not_over_seats CHECK (seats >= available_seats);
