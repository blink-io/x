CREATE TABLE devices (
    id BIGINT NOT NULL DEFAULT nextval('devices_id_seq'::regclass)
    ,name VARCHAR(200) NOT NULL
    ,model VARCHAR(200) NOT NULL
    ,guid VARCHAR(60) NOT NULL
    ,created_at TIMESTAMPTZ NOT NULL
    ,updated_at TIMESTAMPTZ NOT NULL

    ,CONSTRAINT devices_id_pkey PRIMARY KEY (id)
    ,CONSTRAINT devices_guid_key UNIQUE (guid)
);

CREATE TABLE enums (
    id BIGINT NOT NULL DEFAULT nextval('enums_id_seq'::regclass)
    ,status USER_STATUS NOT NULL
    ,created_at TIMESTAMPTZ NOT NULL DEFAULT now()

    ,CONSTRAINT enums_id_pkey PRIMARY KEY (id)
);

CREATE TABLE tags (
    id BIGINT NOT NULL DEFAULT nextval('tags_id_seq'::regclass)
    ,name VARCHAR(60) NOT NULL
    ,code VARCHAR(60) NOT NULL
    ,description TEXT
    ,guid VARCHAR(60) NOT NULL

    ,CONSTRAINT tags_id_pkey PRIMARY KEY (id)
    ,CONSTRAINT tags_guid_key UNIQUE (guid)
);

CREATE TABLE user_devices (
    id BIGINT NOT NULL DEFAULT nextval('user_devices_id_seq'::regclass)
    ,user_id BIGINT NOT NULL
    ,guid VARCHAR(60) NOT NULL
    ,model VARCHAR(200) NOT NULL
    ,name VARCHAR(200) NOT NULL
    ,description TEXT
    ,created_at TIMESTAMPTZ NOT NULL
    ,updated_at TIMESTAMPTZ NOT NULL

    ,CONSTRAINT user_devices_id_pkey PRIMARY KEY (id)
    ,CONSTRAINT user_devices_guid_key UNIQUE (guid)
);

CREATE TABLE users (
    id BIGINT NOT NULL DEFAULT nextval('users_id_seq'::regclass)
    ,guid VARCHAR(60) NOT NULL
    ,username VARCHAR(60) NOT NULL
    ,score DOUBLE PRECISION NOT NULL
    ,created_at TIMESTAMPTZ NOT NULL
    ,updated_at TIMESTAMPTZ NOT NULL

    ,CONSTRAINT users_id_pkey PRIMARY KEY (id)
    ,CONSTRAINT users_guid_key UNIQUE (guid)
    ,CONSTRAINT users_username_key UNIQUE (username)
);
