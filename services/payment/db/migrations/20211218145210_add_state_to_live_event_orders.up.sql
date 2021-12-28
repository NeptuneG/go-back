CREATE TYPE state AS ENUM ('created', 'paid', 'refunded', 'failed');
ALTER TABLE live_event_orders ADD COLUMN state state NOT NULL DEFAULT 'created';
