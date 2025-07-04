package sqtest

import "github.com/blink-io/sq"

type tables struct {
	TblBasic TBL_BASIC
}

var Tables = tables{
	TblBasic: sq.New[TBL_BASIC](""),
}

type TBL_BASIC struct {
	sq.TableStruct
	ID           sq.NumberField `ddl:"type=bigint notnull primarykey default=nextval('tbl_basic_id_seq'::regclass)"`
	N_STR        sq.StringField `ddl:"type=varchar"`
	V_STR        sq.StringField `ddl:"type=varchar notnull"`
	N_ENUM       sq.EnumField   `ddl:"type=user_status"`
	V_ENUM       sq.EnumField   `ddl:"type=user_status notnull"`
	N_INT32      sq.NumberField `ddl:"type=int"`
	V_INT32      sq.NumberField `ddl:"type=int notnull"`
	N_TIME       sq.TimeField   `ddl:"type=timestamptz"`
	V_TIME       sq.TimeField   `ddl:"type=timestamptz notnull"`
	N_UUID       sq.UUIDField   `ddl:"type=uuid"`
	V_UUID       sq.UUIDField   `ddl:"type=uuid notnull"`
	N_JSON       sq.JSONField   `ddl:"type=json"`
	V_JSON       sq.JSONField   `ddl:"type=json notnull"`
	V_STR_ARRAYS sq.ArrayField  `ddl:"type=varchar[] notnull"`
	N_STR_ARRAYS sq.ArrayField  `ddl:"type=varchar[]"`
	N_BYTES      sq.BinaryField `ddl:"type=bytea"`
	V_BYTES      sq.BinaryField `ddl:"type=bytea notnull"`
}

func (t TBL_BASIC) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.ID}
}

func (t TBL_BASIC) PrimaryKeyValues(ID int64) sq.Predicate {
	return t.PrimaryKeys().In(sq.RowValues{{ID}})
}

type EnumTblBasicNEnum string

const EnumTblBasicNEnumUnknown EnumTblBasicNEnum = "unknown"

var EnumTblBasicNEnumValues = []string{
	string(EnumTblBasicNEnumUnknown),
}

func (e EnumTblBasicNEnum) Enumerate() []string {
	return EnumTblBasicNEnumValues
}

type EnumTblBasicVEnum string

const EnumTblBasicVEnumUnknown EnumTblBasicVEnum = "unknown"

var EnumTblBasicVEnumValues = []string{
	string(EnumTblBasicVEnumUnknown),
}

func (e EnumTblBasicVEnum) Enumerate() []string {
	return EnumTblBasicVEnumValues
}
