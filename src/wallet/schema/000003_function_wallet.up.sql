DROP FUNCTION IF EXISTS insert_or_update_wallet() CASCADE;

CREATE OR REPLACE FUNCTION insert_or_update_wallet() RETURNS trigger AS '
    BEGIN
   
    INSERT INTO
        wallet_histories (
            wallet_id,
            order_id,
            last_amount,
            date
        )
    VALUES
        (
            NEW.wallet_id,
            NEW.order_id,
            NEW.last_amount,
            NEW.date
        );

    RETURN NEW;

    END;
' LANGUAGE plpgsql;


CREATE TRIGGER insert_or_update_wallet_trigger
AFTER
INSERT
    OR
UPDATE
    ON "wallets" FOR EACH ROW EXECUTE PROCEDURE insert_or_update_wallet();
