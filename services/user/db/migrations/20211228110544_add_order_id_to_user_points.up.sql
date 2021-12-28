ALTER TABLE "user_points" ADD COLUMN "order_id" uuid;
CREATE UNIQUE INDEX ON "user_points" ("order_id");
