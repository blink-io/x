package tables

import "github.com/blink-io/sq"

type ARRAYS struct {
	sq.TableStruct
	ID          sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('arrays_id_seq'::regclass)"`
	STR_ARRAYS  sq.ArrayField  `ddl:"type=varchar[] notnull"`
	INT4_ARRAYS sq.ArrayField  `ddl:"type=int[] notnull"`
	BOOL_ARRAYS sq.ArrayField  `ddl:"type=boolean[] notnull"`
	CREATED_AT  sq.TimeField   `ddl:"type=timestamptz notnull"`
	V_JSONB     sq.JSONField   `ddl:"type=jsonb"`
	V_JSON      sq.JSONField   `ddl:"type=json"`
	V_UUID      sq.UUIDField   `ddl:"type=uuid"`
}

type DEVICES struct {
	sq.TableStruct
	ID         sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('devices_id_seq'::regclass)"`
	NAME       sq.StringField `ddl:"type=varchar(200) notnull"`
	MODEL      sq.StringField `ddl:"type=varchar(200) notnull"`
	GUID       sq.StringField `ddl:"type=varchar(60) notnull unique"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
	UPDATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
}

type ENUMS struct {
	sq.TableStruct
	ID         sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('enums_id_seq'::regclass)"`
	STATUS     sq.EnumField   `ddl:"type=user_status notnull"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull default=now()"`
}

type HELLO_WORLD struct {
	sq.TableStruct `sq:"public.hello world"`
	ID             sq.NumberField `ddl:"type=bigint notnull primarykey default={nextval('\"hello world_id_seq\"'::regclass)}"`
}

type LOGS struct {
	sq.TableStruct
	ID        sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('logs_id_seq'::regclass)"`
	CONTENT   sq.StringField `ddl:"type=text notnull"`
	CREATD_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
}

type MKEYS struct {
	sq.TableStruct `ddl:"primarykey=id1,id2"`
	ID1            sq.NumberField `ddl:"type=int notnull default=nextval('mkeys_id1_seq'::regclass)"`
	ID2            sq.NumberField `ddl:"type=int notnull default=nextval('mkeys_id2_seq'::regclass)"`
	NAME           sq.StringField `ddl:"type=varchar(60) notnull"`
	CREATED_AT     sq.TimeField   `ddl:"type=timestamptz notnull"`
	GUID           sq.StringField `ddl:"type=varchar(60) notnull unique"`
}

type NEW_WORDS struct {
	sq.TableStruct `sq:"public.new words"`
	ID             sq.NumberField `ddl:"type=bigint notnull primarykey default={nextval('\"new words_id_seq\"'::regclass)}"`
	GUID           sq.StringField `ddl:"type=varchar(60) notnull unique"`
	CREATED_AT     sq.TimeField   `ddl:"type=timestamptz notnull"`
	UPDATED_AT     sq.TimeField   `ddl:"type=timestamptz notnull"`
	CONTENT        sq.StringField `ddl:"type=varchar(200) notnull"`
}

type TAGS struct {
	sq.TableStruct
	ID          sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('tags_id_seq'::regclass)"`
	NAME        sq.StringField `ddl:"type=varchar(60) notnull"`
	CODE        sq.StringField `ddl:"type=varchar(60) notnull"`
	DESCRIPTION sq.StringField `ddl:"type=text"`
	GUID        sq.StringField `ddl:"type=varchar(60) notnull unique"`
	CREATED_AT  sq.TimeField   `ddl:"type=timestamptz notnull"`
}

type USER_DEVICES struct {
	sq.TableStruct
	ID          sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('user_devices_id_seq'::regclass)"`
	USER_ID     sq.NumberField `ddl:"type=bigint notnull"`
	GUID        sq.StringField `ddl:"type=varchar(60) notnull unique"`
	MODEL       sq.StringField `ddl:"type=varchar(200) notnull"`
	NAME        sq.StringField `ddl:"type=varchar(200) notnull"`
	DESCRIPTION sq.StringField `ddl:"type=text"`
	CREATED_AT  sq.TimeField   `ddl:"type=timestamptz notnull"`
	UPDATED_AT  sq.TimeField   `ddl:"type=timestamptz notnull"`
}

type USERS struct {
	sq.TableStruct
	ID         sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('users_id_seq'::regclass)"`
	USERNAME   sq.StringField `ddl:"type=varchar(60) notnull default={''::character varying}"`
	FIRST_NAME sq.StringField `ddl:"type=varchar(60) notnull default={''::character varying}"`
	LAST_NAME  sq.StringField `ddl:"type=varchar(60) notnull default={''::character varying}"`
	LEVEL      sq.NumberField `ddl:"type=smallint notnull default=0"`
	SCORE      sq.NumberField `ddl:"type={double precision} notnull default=0.88"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
	GUID       sq.StringField `ddl:"type=varchar(60) unique"`
	TENANT_ID  sq.NumberField `ddl:"type=bigint notnull default=1"`
	UPDATED_AT sq.TimeField   `ddl:"type=timestamptz notnull default=now()"`
}
