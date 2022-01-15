CREATE TABLE "user_points" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "user_id" uuid NOT NULL,
  "points" integer NOT NULL,
  "description" varchar(255),
  "order_type" varchar(255) NOT NULL,
  "order_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "user_points" ("order_type", "order_id");
