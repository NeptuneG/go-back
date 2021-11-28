CREATE TABLE "live_houses" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "name" varchar NOT NULL,
  "address" varchar,
  "slug" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "live_events" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "live_house_id" uuid NOT NULL,
  "title" varchar NOT NULL,
  "url" varchar NOT NULL,
  "description" text,
  "price_info" varchar,
  "stage_one_open_at" timestamptz,
  "stage_one_start_at" timestamptz NOT NULL,
  "stage_two_open_at" timestamptz,
  "stage_two_start_at" timestamptz,
  "slug" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "artists" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "name" varchar NOT NULL,
  "description" text,
  "slug" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

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

ALTER TABLE "live_events" ADD FOREIGN KEY ("live_house_id") REFERENCES "live_houses" ("id");

ALTER TABLE "user_orders" ADD FOREIGN KEY ("live_event_id") REFERENCES "live_events" ("id");

ALTER TABLE "user_orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "live_houses" ("name");

CREATE UNIQUE INDEX ON "live_houses" ("slug");

CREATE INDEX ON "live_events" ("live_house_id");

CREATE INDEX ON "live_events" ("title");

CREATE UNIQUE INDEX ON "live_events" ("slug");

CREATE UNIQUE INDEX ON "live_events" ("url");

CREATE INDEX ON "artists" ("name");

CREATE UNIQUE INDEX ON "artists" ("slug");

CREATE UNIQUE INDEX ON "users" ("email");

CREATE UNIQUE INDEX ON "users" ("reset_password_token");
