CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE SCHEMA erpv1 AUTHORIZATION erp;
CREATE TABLE IF NOT EXISTS erpv1."users"
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username character varying COLLATE pg_catalog."default" NOT NULL,
    phone_number character varying COLLATE pg_catalog."default",
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT username UNIQUE (username)
);

-- erpv1.emails
CREATE SEQUENCE IF NOT EXISTS erpv1.emails_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;
ALTER SEQUENCE erpv1.emails_id_seq
    OWNER TO erp;
CREATE TABLE IF NOT EXISTS erpv1.emails
(
    id bigint NOT NULL DEFAULT nextval('erpv1.emails_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    user_id uuid,
    address character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT emails_pkey PRIMARY KEY (id),
    CONSTRAINT fk_erpv1_users_emails FOREIGN KEY (user_id)
        REFERENCES erpv1.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);


-- erpv1.auths
CREATE SEQUENCE IF NOT EXISTS erpv1.auths_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;
ALTER SEQUENCE erpv1.auths_id_seq
    OWNER TO erp;
CREATE TABLE IF NOT EXISTS erpv1.auths
(
    id bigint NOT NULL DEFAULT nextval('erpv1.auths_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    user_id uuid,
    password character varying(64) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT auths_pkey PRIMARY KEY (id),
    CONSTRAINT fk_erpv1_users_auth FOREIGN KEY (user_id)
        REFERENCES erpv1.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

ALTER TABLE IF EXISTS erpv1.users
    OWNER to erp;
ALTER TABLE IF EXISTS erpv1.emails
    OWNER to erp;
ALTER TABLE IF EXISTS erpv1.auths
    OWNER to erp;

CREATE INDEX IF NOT EXISTS idx_erpv1_users_deleted_at
    ON erpv1.users USING btree
    (deleted_at ASC NULLS LAST);
CREATE INDEX IF NOT EXISTS idx_erpv1_emails_deleted_at
    ON erpv1.emails USING btree
    (deleted_at ASC NULLS LAST);
CREATE INDEX IF NOT EXISTS idx_erpv1_auths_deleted_at
    ON erpv1.auths USING btree
    (deleted_at ASC NULLS LAST);
