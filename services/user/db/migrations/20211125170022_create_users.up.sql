CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "email" varchar NOT NULL,
  "encrypted_password" varchar NOT NULL,
  "reset_password_token" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_orders" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "live_event_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "user_orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE UNIQUE INDEX ON "users" ("email");

CREATE UNIQUE INDEX ON "users" ("reset_password_token");
