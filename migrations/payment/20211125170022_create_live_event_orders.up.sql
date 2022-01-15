CREATE TABLE "live_event_orders" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "user_id" uuid NOT NULL,
  "live_event_id" uuid NOT NULL,
  "price" integer NOT NULL,
  "user_points" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
