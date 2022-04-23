CREATE SEQUENCE IF NOT EXISTS product_seq;

CREATE TABLE "products" (
  "product_id" BIGINT DEFAULT nextval('product_seq') PRIMARY KEY,
  "price" BIGINT NOT NULL,
  "name" VARCHAR NOT NULL
);

CREATE INDEX ON "products" ("product_id");

INSERT into "products" (
    "price",
    "name") 
  VALUES (
    10000,
    'simpati 10000'
  );

INSERT into "products" (
    "price",
    "name") 
  VALUES (
    25000,
    'simpati 25000'
  );
