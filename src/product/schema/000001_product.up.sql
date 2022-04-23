CREATE SEQUENCE IF NOT EXISTS product_seq;

CREATE TABLE "products" (
  "product_id" BIGINT DEFAULT nextval('product_seq') PRIMARY KEY,
  "price" BIGINT NOT NULL,
  "name" VARCHAR NOT NULL
);

CREATE INDEX ON "products" ("product_id");