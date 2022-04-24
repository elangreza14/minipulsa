CREATE SEQUENCE IF NOT EXISTS order_seq;

  -- int64 order_id = 1;
  -- int64 product_id = 2;
  -- int64 user_id = 3;
  -- int64 wallet_history_id = 4;
  -- string status = 5;
  -- string date = 6;

CREATE TABLE "orders" (
  "order_id" BIGINT DEFAULT nextval('order_seq') PRIMARY KEY,
  "product_id" BIGINT NOT NULL,
  "user_id" BIGINT NOT NULL,
  "wallet_history_id" BIGINT,
  "status" VARCHAR NOT NULL,
  "date" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON "orders" ("order_id");