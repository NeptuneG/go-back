ALTER TABLE live_events ADD COLUMN available_seats INTEGER
DEFAULT 100 NOT NULL CONSTRAINT positive CHECK (available_seats >= 0);
