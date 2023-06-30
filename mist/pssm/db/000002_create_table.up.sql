CREATE TABLE IF NOT EXISTS passmexample."commodity"
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name character varying NOT NULL,
    price numeric NOT NULL DEFAULT 0,
    stock bigint,
    category character varying,
    image_url character varying,
    unit character varying,
    barcode character varying,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id),
    CONSTRAINT name UNIQUE (name)
);

CREATE INDEX IF NOT EXISTS idx_passmexample_commodity_deleted_at
    ON passmexample.commodity USING btree
    (deleted_at ASC NULLS LAST);
CREATE INDEX idx_passmexample_commodity_category
    ON passmexample.commodity USING btree
    (category ASC NULLS LAST);

ALTER TABLE IF EXISTS passmexample.commodity
    OWNER to erp;