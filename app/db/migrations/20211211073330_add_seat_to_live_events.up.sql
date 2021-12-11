ALTER TABLE live_events ADD COLUMN available_seats INTEGER
DEFAULT 100 CONSTRAINT positive CHECK (available_seats >= 0);
