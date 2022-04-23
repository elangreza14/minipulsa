DROP FUNCTION IF EXISTS insert_or_update_order() CASCADE;

CREATE OR REPLACE FUNCTION insert_or_update_order() RETURNS trigger AS '
    BEGIN
   
    INSERT INTO
        order_histories (
            order_id,
            status,
            date
        )
    VALUES
        (
            NEW.order_id,
            NEW.status,
            NEW.date
        );

    RETURN NEW;

    END;
' LANGUAGE plpgsql;


CREATE TRIGGER insert_or_update_order_trigger
AFTER
INSERT
    OR
UPDATE
    ON "orders" FOR EACH ROW EXECUTE PROCEDURE insert_or_update_order();
