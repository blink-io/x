package orm

import "github.com/bokwoon95/sq"

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

type TAGS struct {
	sq.TableStruct
	ID          sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('tags_id_seq'::regclass)"`
	NAME        sq.StringField `ddl:"type=varchar(60) notnull"`
	CODE        sq.StringField `ddl:"type=varchar(60) notnull"`
	DESCRIPTION sq.StringField `ddl:"type=text"`
	CREATED_AT  sq.TimeField   `ddl:"type=timestamptz notnull"`
	GUID        sq.StringField `ddl:"type=varchar(60) notnull unique"`
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
	GUID       sq.StringField `ddl:"type=varchar(60) notnull unique"`
	USERNAME   sq.StringField `ddl:"type=varchar(60) notnull unique"`
	SCORE      sq.NumberField `ddl:"type={double precision} notnull"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
	UPDATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
}

type TVALS struct {
	sq.TableStruct `ddl:"primarykey=iid,sid"`
	IID            sq.NumberField `ddl:"type=int notnull"`
	SID            sq.StringField `ddl:"type=int notnull"`
	NAME           sq.StringField `ddl:"type=varchar(60) notnull"`
	CREATED_AT     sq.TimeField   `ddl:"type=timestamptz notnull"`
}
