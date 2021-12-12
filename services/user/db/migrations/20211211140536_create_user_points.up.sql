CREATE TABLE "user_points" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "user_id" uuid NOT NULL,
  "points" integer NOT NULL,
  "description" varchar(255),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "user_points" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
