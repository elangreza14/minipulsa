CREATE USER minipulsa WITH PASSWORD 'minipulsa' CREATEDB;

CREATE DATABASE "authentication"
    WITH 
    OWNER = minipulsa
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE DATABASE "product"
    WITH 
    OWNER = minipulsa
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE DATABASE "order"
    WITH 
    OWNER = minipulsa
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

CREATE DATABASE "wallet"
    WITH 
    OWNER = minipulsa
    ENCODING = 'UTF8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;