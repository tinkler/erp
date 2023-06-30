CREATE TABLE passmexample.store_in
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id)
);
ALTER TABLE IF EXISTS passmexample.store_in
    OWNER to erp;


CREATE SEQUENCE IF NOT EXISTS passmexample.store_in_item_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;
ALTER SEQUENCE passmexample.store_in_item_id_seq
    OWNER TO erp;
CREATE TABLE IF NOT EXISTS passmexample.store_in_item
(
    id bigint NOT NULL DEFAULT nextval('passmexample.store_in_item_id_seq'::regclass),
    index smallint NOT NULL DEFAULT 0,
    store_in_id uuid,
    commodity_id uuid,
    price numeric NOT NULL DEFAULT 0,
    quantity bigint,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id),
    CONSTRAINT fk_passmexample_store_in_item_store_in_id FOREIGN KEY (store_in_id)
        REFERENCES passmexample.store_in (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
ALTER TABLE IF EXISTS passmexample.store_in_item
    OWNER to erp;