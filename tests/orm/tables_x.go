package orm

import (
	"context"
	"time"

	sqx "github.com/blink-io/x/sql/builder/sq"
	"github.com/bokwoon95/sq"
)

type tables struct {
	Tags    TAGS
	TagsBak TAGS_BAK
	Users   USERS
	Tvals   TVALS
	Mkeys   MKEYS
	Arrays  ARRAYS
}

var Tables = tables{
	Tags:    sq.New[TAGS](""),
	TagsBak: sq.New[TAGS_BAK](""),
	Users:   sq.New[USERS](""),
	Tvals:   sq.New[TVALS](""),
	Mkeys:   sq.New[MKEYS](""),
	Arrays:  sq.New[ARRAYS](""),
}

var _ sq.PolicyTable = Tables.Users

const (
	defaultTenantID = 1
)

func (t TVALS) PrimaryKeys() sq.RowValue {
	return sq.RowValue{t.IID, t.SID}
}

func (t TVALS) PrimaryKeyValues(iid int64, sid string) sq.Predicate {
	return t.PrimaryKeys().Eq(sq.RowValues{{iid, sid}})
}

func (tbl USERS) Policy(ctx context.Context, dialect string) (sq.Predicate, error) {
	tenantID, ok := ctx.Value(tbl.TENANT_ID.GetName()).(int)
	if !ok {
		return sq.And(sqx.AlwaysTrueExpr, tbl.TENANT_ID.EqInt(defaultTenantID)), nil
	}
	return tbl.TENANT_ID.EqInt(tenantID), nil
}

func (tbl USERS) InsertMapper(rs ...User) func(c *sq.Column) {
	return func(c *sq.Column) {
		for _, r := range rs {
			if r.ID > 0 {
				c.SetInt(tbl.ID, r.ID)
			}
			c.SetString(tbl.GUID, r.GUID)
			c.SetString(tbl.USERNAME, r.Username)
			c.SetString(tbl.FIRST_NAME, r.FirstName)
			c.SetString(tbl.LAST_NAME, r.LastName)
			c.SetFloat64(tbl.SCORE, r.Score)
			c.SetInt(tbl.LEVEL, r.Level)
			c.SetTime(tbl.CREATED_AT, r.CreatedAt)
			c.SetTime(tbl.UPDATED_AT, r.UpdatedAt)
			c.SetInt(tbl.TENANT_ID, r.TenantID)
		}
	}
}

func (tbl USERS) QueryMapper() func(*sq.Row) *User {
	return func(r *sq.Row) *User {
		u := &User{
			ID:        r.IntField(tbl.ID),
			GUID:      r.StringField(tbl.GUID),
			Username:  r.StringField(tbl.USERNAME),
			FirstName: r.StringField(tbl.FIRST_NAME),
			LastName:  r.StringField(tbl.LAST_NAME),
			Score:     r.Float64Field(tbl.SCORE),
			Level:     r.IntField(tbl.LEVEL),
			TenantID:  r.IntField(tbl.TENANT_ID),
		}

		dd := sq.DefaultDialect.Load()
		dz := sq.DialectSQLite
		if dz == *dd {
			crstr := r.String("created_at")
			upstr := r.String("updated_at")
			if ct, err := time.Parse(time.RFC3339Nano, crstr); err == nil {
				u.CreatedAt = ct
			}
			if ct, err := time.Parse(time.RFC3339Nano, upstr); err == nil {
				u.UpdatedAt = ct
			}
		} else {
			u.CreatedAt = r.TimeField(tbl.CREATED_AT)
			u.UpdatedAt = r.TimeField(tbl.UPDATED_AT)
		}

		return u
	}
}
