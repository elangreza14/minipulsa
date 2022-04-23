CREATE SEQUENCE IF NOT EXISTS user_seq;

CREATE TABLE "users" (
  "user_id" BIGINT DEFAULT nextval('user_seq') PRIMARY KEY,
  "email" VARCHAR NOT NULL UNIQUE,
  "password" VARCHAR NOT NULL,
  "token" VARCHAR NOT NULL
);

CREATE INDEX ON "users" ("user_id");