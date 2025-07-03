package orm

import "github.com/blink-io/sq"

type tables struct {
	UserCredentials USER_CREDENTIALS
	Arrays          ARRAYS
	Devices         DEVICES
	Enums           ENUMS
	HelloWorld      HELLO_WORLD
	Jsons           JSONS
	Logs            LOGS
	Mkeys           MKEYS
	NewWords        NEW_WORDS
	TJsons          T_JSONS
	Tags            TAGS
	TagsBak         TAGS_BAK
	TsArrays        TS_ARRAYS
	UserDevices     USER_DEVICES
	Users           USERS
}

var Tables = tables{
	UserCredentials: sq.New[USER_CREDENTIALS](""),
	Arrays:          sq.New[ARRAYS](""),
	Devices:         sq.New[DEVICES](""),
	Enums:           sq.New[ENUMS](""),
	HelloWorld:      sq.New[HELLO_WORLD](""),
	Jsons:           sq.New[JSONS](""),
	Logs:            sq.New[LOGS](""),
	Mkeys:           sq.New[MKEYS](""),
	NewWords:        sq.New[NEW_WORDS](""),
	TJsons:          sq.New[T_JSONS](""),
	Tags:            sq.New[TAGS](""),
	TagsBak:         sq.New[TAGS_BAK](""),
	TsArrays:        sq.New[TS_ARRAYS](""),
	UserDevices:     sq.New[USER_DEVICES](""),
	Users:           sq.New[USERS](""),
}

type USER_CREDENTIALS struct {
	sq.TableStruct `sq:"iam.user_credentials"`
	ID             sq.NumberField  `ddl:"type=bigint notnull primarykey default=nextval('iam.user_credentials_id_seq'::regclass)"`
	GUID           sq.StringField  `ddl:"type=varchar(60)"`
	CREATED_AT     sq.TimeField    `ddl:"type=timestamptz(6) notnull"`
	UPDATED_AT     sq.TimeField    `ddl:"type=timestamptz(6) notnull"`
	CREATED_BY     sq.StringField  `ddl:"type=varchar(60)"`
	UPDATED_BY     sq.StringField  `ddl:"type=varchar(60)"`
	DELETED_AT     sq.TimeField    `ddl:"type=timestamptz(6)"`
	DELETED_BY     sq.StringField  `ddl:"type=varchar(60)"`
	IS_DELETED     sq.BooleanField `ddl:"type=boolean"`
	USER_ID        sq.StringField  `ddl:"type=varchar notnull"`
	STATUS         sq.StringField  `ddl:"type=varchar notnull"`
	TYPE           sq.StringField  `ddl:"type=varchar notnull"`
	PAYLOAD        sq.StringField  `ddl:"type=varchar notnull"`
	SCHEME         sq.StringField  `ddl:"type=varchar notnull"`
	FACTOR         sq.StringField  `ddl:"type=varchar notnull"`
	EXPIRY         sq.NumberField  `ddl:"type=bigint notnull default='-1'::integer"`
	DESCRIPTION    sq.StringField  `ddl:"type=text"`
}

func (t USER_CREDENTIALS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t USER_CREDENTIALS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}

type ARRAYS struct {
	sq.TableStruct
	ID            sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('arrays_id_seq'::regclass)"`
	STR_ARRAYS    sq.ArrayField  `ddl:"type=varchar[] notnull"`
	INT4_ARRAYS   sq.ArrayField  `ddl:"type=int[] notnull"`
	BOOL_ARRAYS   sq.ArrayField  `ddl:"type=boolean[] notnull"`
	CREATED_AT    sq.TimeField   `ddl:"type=timestamptz notnull"`
	V_JSONB       sq.JSONField   `ddl:"type=jsonb"`
	V_JSON        sq.JSONField   `ddl:"type=json"`
	V_UUID        sq.UUIDField   `ddl:"type=uuid"`
	JSONB_ARRAYS  sq.ArrayField  `ddl:"type=jsonb[]"`
	JSON_ARRAYS   sq.ArrayField  `ddl:"type=json[]"`
	UUID_ARRAYS   sq.ArrayField  `ddl:"type=uuid[]"`
	INT_AAA       sq.ArrayField  `ddl:"type=int[]"`
	TS_ARRAYS     sq.ArrayField  `ddl:"type=timestamptz[]"`
	INT2_ARRAYS   sq.ArrayField  `ddl:"type=smallint[]"`
	REMARK        sq.StringField `ddl:"type=varchar"`
	STATUS_ARRAYS sq.ArrayField  `ddl:"type=user_status[]"`
}

func (t ARRAYS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t ARRAYS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t DEVICES) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t ENUMS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}

type EnumEnumsStatus string

const EnumEnumsStatusUnknown EnumEnumsStatus = "unknown"

var EnumEnumsStatusValues = []string{
	string(EnumEnumsStatusUnknown),
}

func (e EnumEnumsStatus) Enumerate() []string {
	return EnumEnumsStatusValues
}

type EnumEnumsMoodx string

const EnumEnumsMoodxUnknown EnumEnumsMoodx = "unknown"

var EnumEnumsMoodxValues = []string{
	string(EnumEnumsMoodxUnknown),
}

func (e EnumEnumsMoodx) Enumerate() []string {
	return EnumEnumsMoodxValues
}

type HELLO_WORLD struct {
	sq.TableStruct
	ID sq.NumberField `ddl:"type=bigint notnull primarykey default={nextval('\"hello world_id_seq\"'::regclass)}"`
}

func (t HELLO_WORLD) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t HELLO_WORLD) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}

type JSONS struct {
	sq.TableStruct
	ID     sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('jsons_id_seq'::regclass)"`
	V_JSON sq.StringField `ddl:"type=varchar"`
	V_UUID sq.UUIDField   `ddl:"type=uuid"`
}

func (t JSONS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t JSONS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t LOGS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t MKEYS) PrimaryKeyValues(Id1 int32, Id2 int32) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{Id1, Id2}})
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

func (t NEW_WORDS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}

type T_JSONS struct {
	sq.TableStruct
	ID      sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('t_jsons_id_seq'::regclass)"`
	V_JSON  sq.JSONField   `ddl:"type=json"`
	V_JSONB sq.JSONField   `ddl:"type=jsonb"`
	A_JSON  sq.ArrayField  `ddl:"type=json[]"`
}

func (t T_JSONS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t T_JSONS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t TAGS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t TAGS_BAK) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}

type TS_ARRAYS struct {
	sq.TableStruct
	ID  sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('ts_arrays_id_seq'::regclass)"`
	TSA sq.ArrayField  `ddl:"type=timestamptz[]"`
}

func (t TS_ARRAYS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TS_ARRAYS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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

func (t USER_DEVICES) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
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
	GUID       sq.StringField `ddl:"type=varchar(60) notnull unique"`
	TENANT_ID  sq.NumberField `ddl:"type=bigint notnull default=1"`
	UPDATED_AT sq.TimeField   `ddl:"type=timestamptz notnull default=now()"`
}

func (t USERS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t USERS) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}
