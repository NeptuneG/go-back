CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (public.gen_random_uuid()),
  "email" varchar NOT NULL,
  "encrypted_password" varchar NOT NULL,
  "reset_password_token" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "users" ("email");

CREATE UNIQUE INDEX ON "users" ("reset_password_token");
