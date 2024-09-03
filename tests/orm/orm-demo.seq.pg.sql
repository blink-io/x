-- public.devices_id_seq definition

-- DROP SEQUENCE public.devices_id_seq;

CREATE SEQUENCE public.devices_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;


-- public.enums_id_seq definition

-- DROP SEQUENCE public.enums_id_seq;

CREATE SEQUENCE public.enums_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;


-- public.tags_id_seq definition

-- DROP SEQUENCE public.tags_id_seq;

CREATE SEQUENCE public.tags_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;


-- public.user_devices_id_seq definition

-- DROP SEQUENCE public.user_devices_id_seq;

CREATE SEQUENCE public.user_devices_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;


-- public.users_id_seq definition

-- DROP SEQUENCE public.users_id_seq;

CREATE SEQUENCE public.users_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
	CACHE 1
	NO CYCLE;