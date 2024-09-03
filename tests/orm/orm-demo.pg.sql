-- DROP TYPE public.user_status;

CREATE TYPE public.user_status AS ENUM (
    'active',
    'blocked');

-- public.devices definition

-- Drop table

-- DROP TABLE public.devices;

CREATE TABLE public.devices
(
    id         bigserial    NOT NULL,
    "name"     varchar(200) NOT NULL,
    model      varchar(200) NOT NULL,
    guid       varchar(60)  NOT NULL,
    created_at timestamptz  NOT NULL,
    updated_at timestamptz  NOT NULL,
    CONSTRAINT devices_pk_id PRIMARY KEY (id),
    CONSTRAINT devices_ukey_guid UNIQUE (guid)
);


-- public.enums definition

-- Drop table

-- DROP TABLE public.enums;

CREATE TABLE public.enums
(
    id         bigserial                 NOT NULL,
    status public.user_status NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL,
    CONSTRAINT enums_pk_id PRIMARY KEY (id)
);


-- public.sqddl_history definition

-- Drop table

-- DROP TABLE public.sqddl_history;

CREATE TABLE public.sqddl_history
(
    filename      varchar(255) NOT NULL,
    checksum      varchar(64) NULL,
    started_at    timestamptz NULL,
    time_taken_ns int8 NULL,
    success       bool NULL,
    CONSTRAINT sqddl_filename_pkey PRIMARY KEY (filename)
);


-- public.tags definition

-- Drop table

-- DROP TABLE public.tags;

CREATE TABLE public.tags
(
    id          bigserial   NOT NULL,
    "name"      varchar(60) NOT NULL,
    code        varchar(60) NOT NULL,
    description text NULL,
    guid        varchar(60) NOT NULL,
    CONSTRAINT tags_pk_id PRIMARY KEY (id),
    CONSTRAINT tags_ukey_guid UNIQUE (guid)
);


-- public.user_devices definition

-- Drop table

-- DROP TABLE public.user_devices;

CREATE TABLE public.user_devices
(
    id          bigserial    NOT NULL,
    user_id     int8         NOT NULL,
    guid        varchar(60)  NOT NULL,
    model       varchar(200) NOT NULL,
    "name"      varchar(200) NOT NULL,
    description text NULL,
    created_at  timestamptz  NOT NULL,
    updated_at  timestamptz  NOT NULL,
    CONSTRAINT user_devices_pk_id PRIMARY KEY (id),
    CONSTRAINT user_devices_ukey_guid UNIQUE (guid)
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    id         bigserial   NOT NULL,
    guid       varchar(60) NOT NULL,
    username   varchar(60) NOT NULL,
    score      float8      NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT users_pk_id PRIMARY KEY (id),
    CONSTRAINT users_ukey_guid UNIQUE (guid),
    CONSTRAINT users_ukey_username UNIQUE (username)
);