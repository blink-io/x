package orm

import "github.com/blink-io/sq"

type tables struct {
	Arrays      ARRAYS
	Devices     DEVICES
	Enums       ENUMS
	HelloWorld  HELLO_WORLD
	Logs        LOGS
	Mkeys       MKEYS
	NewWords    NEW_WORDS
	Tags        TAGS
	TagsBak     TAGS_BAK
	UserDevices USER_DEVICES
	Users       USERS
}

var Tables = tables{
	Arrays:      sq.New[ARRAYS](""),
	Devices:     sq.New[DEVICES](""),
	Enums:       sq.New[ENUMS](""),
	HelloWorld:  sq.New[HELLO_WORLD](""),
	Logs:        sq.New[LOGS](""),
	Mkeys:       sq.New[MKEYS](""),
	NewWords:    sq.New[NEW_WORDS](""),
	Tags:        sq.New[TAGS](""),
	TagsBak:     sq.New[TAGS_BAK](""),
	UserDevices: sq.New[USER_DEVICES](""),
	Users:       sq.New[USERS](""),
}

type ARRAYS struct {
	sq.TableStruct
	ID           sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('arrays_id_seq'::regclass)"`
	STR_ARRAYS   sq.ArrayField  `ddl:"type=varchar[] notnull"`
	INT4_ARRAYS  sq.ArrayField  `ddl:"type=int[] notnull"`
	BOOL_ARRAYS  sq.ArrayField  `ddl:"type=boolean[] notnull"`
	CREATED_AT   sq.TimeField   `ddl:"type=timestamptz notnull"`
	V_JSONB      sq.JSONField   `ddl:"type=jsonb"`
	V_JSON       sq.JSONField   `ddl:"type=json"`
	V_UUID       sq.UUIDField   `ddl:"type=uuid"`
	JSONB_ARRAYS sq.ArrayField  `ddl:"type=jsonb[]"`
	JSON_ARRAYS  sq.ArrayField  `ddl:"type=json[]"`
	UUID_ARRAYS  sq.ArrayField  `ddl:"type=uuid[]"`
}

func (t ARRAYS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t ARRAYS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
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

func (t DEVICES) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t DEVICES) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
}

type ENUMS struct {
	sq.TableStruct
	ID         sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('enums_id_seq'::regclass)"`
	STATUS     sq.EnumField   `ddl:"type=user_status notnull"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull default=now()"`
	MOODX      sq.EnumField   `ddl:"type=mood"`
}

func (t ENUMS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t ENUMS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
}

type EnumEnumsStatus string

func (e EnumEnumsStatus) Enumerate() []string {
	//TODO Add more
	return []string{}
}

type EnumEnumsMoodx string

func (e EnumEnumsMoodx) Enumerate() []string {
	//TODO Add more
	return []string{}
}

type HELLO_WORLD struct {
	sq.TableStruct
	ID sq.NumberField `ddl:"type=bigint notnull primarykey default={nextval('\"hello world_id_seq\"'::regclass)}"`
}

func (t HELLO_WORLD) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t HELLO_WORLD) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
}

type LOGS struct {
	sq.TableStruct
	ID         sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('logs_id_seq'::regclass)"`
	CONTENT    sq.StringField `ddl:"type=text notnull"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
	AUDITED_AT sq.TimeField   `ddl:"type=timestamptz"`
}

func (t LOGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t LOGS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
}

type MKEYS struct {
	sq.TableStruct `ddl:"primarykey=id1,id2"`
	ID1            sq.NumberField `ddl:"type=int notnull default=nextval('mkeys_id1_seq'::regclass)"`
	ID2            sq.NumberField `ddl:"type=int notnull default=nextval('mkeys_id2_seq'::regclass)"`
	NAME           sq.StringField `ddl:"type=varchar(60) notnull"`
	CREATED_AT     sq.TimeField   `ddl:"type=timestamptz notnull"`
	GUID           sq.StringField `ddl:"type=varchar(60) notnull unique"`
}

func (t MKEYS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID1, t.ID2}
}

func (t MKEYS) PrimaryKeyValues(id1 int32, id2 int32) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id1, id2}})
}

type NEW_WORDS struct {
	sq.TableStruct
	ID         sq.NumberField `ddl:"type=bigint notnull primarykey default={nextval('\"new words_id_seq\"'::regclass)}"`
	GUID       sq.StringField `ddl:"type=varchar(60) notnull unique"`
	CREATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
	UPDATED_AT sq.TimeField   `ddl:"type=timestamptz notnull"`
	CONTENT    sq.StringField `ddl:"type=varchar(200) notnull"`
}

func (t NEW_WORDS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t NEW_WORDS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
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

func (t TAGS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TAGS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
}

type TAGS_BAK struct {
	sq.TableStruct
	ID          sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('tags_bak_id_seq'::regclass)"`
	NAME        sq.StringField `ddl:"type=varchar(60) notnull"`
	CODE        sq.StringField `ddl:"type=varchar(60) notnull"`
	DESCRIPTION sq.StringField `ddl:"type=text"`
	GUID        sq.StringField `ddl:"type=varchar(60) notnull unique"`
	CREATED_AT  sq.TimeField   `ddl:"type=timestamptz notnull"`
}

func (t TAGS_BAK) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TAGS_BAK) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
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

func (t USER_DEVICES) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t USER_DEVICES) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
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

func (t USERS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t USERS) PrimaryKeyValues(id int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{id}})
}
