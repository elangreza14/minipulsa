CREATE SEQUENCE IF NOT EXISTS order_history_seq;

CREATE TABLE "order_histories" (
  "order_history_id" BIGINT DEFAULT nextval('order_history_seq') PRIMARY KEY,
  "order_id" BIGINT NOT NULL,
  "status" VARCHAR NOT NULL,
  "date" TIMESTAMPTZ NOT NULL
);

CREATE INDEX ON "order_histories" ("order_id");

ALTER TABLE "order_histories" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id") ON DELETE CASCADE;
