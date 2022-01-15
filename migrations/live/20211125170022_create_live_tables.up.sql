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

ALTER TABLE "live_events" ADD FOREIGN KEY ("live_house_id") REFERENCES "live_houses" ("id");

CREATE INDEX ON "live_houses" ("name");

CREATE UNIQUE INDEX ON "live_houses" ("slug");

CREATE INDEX ON "live_events" ("live_house_id");

CREATE INDEX ON "live_events" ("title");

CREATE UNIQUE INDEX ON "live_events" ("slug");

CREATE UNIQUE INDEX ON "live_events" ("url");
