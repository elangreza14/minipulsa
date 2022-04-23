CREATE SEQUENCE IF NOT EXISTS order_seq;

CREATE TABLE "orders" (
  "order_id" BIGINT DEFAULT nextval('order_seq') PRIMARY KEY,
  "product_id" BIGINT NOT NULL,
  "user_id" BIGINT NOT NULL,
  "status" VARCHAR NOT NULL,
  "date" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON "orders" ("order_id");