CREATE DATABASE abishar_bpjs_test;

CREATE TABLE requests (
    request_id integer NOT NULL,
    created_at timestamp without time zone
);

CREATE TABLE request_data (
    id integer NOT NULL,
    customer character varying(50) NOT NULL,
    quantity integer,
    price numeric(10,2),
    "timestamp" timestamp without time zone,
    request_id integer
);