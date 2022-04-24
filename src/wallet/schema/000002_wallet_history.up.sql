CREATE SEQUENCE IF NOT EXISTS wallet_history_seq;

CREATE TABLE "wallet_histories" (
  "wallet_history_id" BIGINT DEFAULT nextval('wallet_history_seq') PRIMARY KEY,
  "wallet_id" BIGINT NOT NULL,
  "order_id" BIGINT,
  "last_amount" BIGINT NOT NULL,
  "date" TIMESTAMPTZ NOT NULL
);

CREATE INDEX ON "wallet_histories" ("wallet_id");

ALTER TABLE "wallet_histories" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("wallet_id") ON DELETE CASCADE;
