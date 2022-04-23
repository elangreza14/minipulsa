CREATE SEQUENCE IF NOT EXISTS wallet_seq;

CREATE TABLE "wallets" (
  "wallet_id" BIGINT DEFAULT nextval('wallet_seq') PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "amount" BIGINT NOT NULL,
  "date" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON "wallets" ("wallet_id");