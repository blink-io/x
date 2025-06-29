package orm

import (
	"context"
	"errors"
	"fmt"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/blink-io/sq"
	"github.com/blink-io/x/log/slog/handlers/color"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sanity-io/litter"
)

var _log = slog.New(color.New(os.Stderr, color.SetLevel(slog.LevelDebug)))

var ctx = context.Background()

func init() {
	l := sq.NewLogger(os.Stdout, "", log.LstdFlags, sq.LoggerConfig{
		ShowCaller:    true,
		ShowTimeTaken: true,
		HideArgs:      true,
	})
	sq.SetDefaultLogQuery(func(ctx context.Context, stats sq.QueryStats) {
		l.LogQuery(ctx, stats)
	})
}

var _ fmt.Stringer = (*User)(nil)
var _ fmt.Stringer = (*UserDevice)(nil)

var (
	UserTable       = sq.New[USERS]("u1")
	UserDeviceTable = sq.New[USER_DEVICES]("u2")
	DeviceTable     = sq.New[DEVICES]("u3")
	TagTable        = sq.New[TAGS]("u4")
)

type UserStatus string

const (
	UserStatusActive  EnumEnumsStatus = "active"
	UserStatusBlocked EnumEnumsStatus = "blocked"
)

func (v UserDevice) String() string {
	return litter.Sdump(v)
}

func (v UserStatus) String() string {
	return string(v)
}

func (m User) String() string {
	return litter.Sdump(m)
}

type UserWithDevice struct {
	UserID      int64  `db:"user_id"`
	UserGUID    string `db:"user_guid"`
	Username    string `db:"username"`
	DeviceGUID  string `db:"device_guid"`
	DeviceName  string `db:"device_name"`
	DeviceModel string `db:"device_model"`
}

func (m UserWithDevice) String() string {
	return litter.Sdump(m)
}

func (m UserSetter) Overwrite(r *User) {
	if !m.ID.IsUnset() {
		r.ID, _ = m.ID.Get()
	}
	if !m.GUID.IsUnset() {
		r.GUID, _ = m.GUID.Get()
	}
	if !m.Username.IsUnset() {
		r.Username, _ = m.Username.Get()
	}
	if !m.Score.IsUnset() {
		r.Score, _ = m.Score.Get()
	}
	if !m.CreatedAt.IsUnset() {
		r.CreatedAt, _ = m.CreatedAt.Get()
	}
	if !m.UpdatedAt.IsUnset() {
		r.CreatedAt, _ = m.CreatedAt.Get()
	}
}

func (m UserSetter) Update(db sq.DB) error {
	tbl := UserTable
	if id, ok := m.ID.Get(); ok {
		_, err := sq.Exec(
			db,
			sq.Update(tbl).
				SetFunc(m.SetColumns).
				Where(tbl.ID.EqInt64(id)),
		)
		return err
	} else {
		return errors.New("id is required")
	}
}

func (m UserSetter) setColumns(c *sq.Column, withID bool) {
	tbl := UserTable
	if withID && !m.ID.IsUnset() {
		v, _ := m.ID.Get()
		c.SetInt64(tbl.ID, v)
	}
	if !m.GUID.IsUnset() {
		v, _ := m.GUID.Get()
		c.SetString(tbl.GUID, v)
	}
	if !m.Username.IsUnset() {
		v, _ := m.Username.Get()
		c.SetString(tbl.USERNAME, v)
	}
	if !m.Score.IsUnset() {
		v, _ := m.Score.Get()
		c.SetFloat64(tbl.SCORE, v)
	}
	if !m.CreatedAt.IsUnset() {
		v, _ := m.CreatedAt.Get()
		c.SetTime(tbl.CREATED_AT, v)
	}
	if !m.UpdatedAt.IsUnset() {
		v, _ := m.UpdatedAt.Get()
		c.SetTime(tbl.UPDATED_AT, v)
	}
}

func (m UserSetter) SetColumns(ctx context.Context, c *sq.Column) {
	m.setColumns(c, true)
}

func (m Device) Insert(db sq.DB) error {
	tbl := DeviceTable
	_, err := sq.Exec(db, sq.InsertInto(tbl).ColumnValues(func(ctx context.Context, c *sq.Column) {
		c.SetString(tbl.GUID, m.GUID)
		c.SetString(tbl.NAME, m.Name)
		c.SetString(tbl.MODEL, m.Model)
		c.SetTime(tbl.CREATED_AT, m.CreatedAt)
		c.SetTime(tbl.UPDATED_AT, m.UpdatedAt)
	}))
	return err
}

func (m Device) String() string {
	return litter.Sdump(m)
}

type Model struct {
	Dialect string `db:"dialect"`
	Name    string `db:"name"`
	Version string `db:"version"`
	Current string `db:"current"`
}

func (m Model) WithName(name string) Model {
	m.Name = name
	return m
}

func (m Model) WithVersion(version string) Model {
	m.Version = version
	return m
}

func (m Model) String() string {
	return litter.Sdump(m)
}

func userJoinDeviceMapRowMapper() func(*sq.Row) map[string]any {
	return func(r *sq.Row) map[string]any {
		//t := UserTable
		//joinTbl := UserDeviceTable
		mm := make(map[string]any)
		mm["id"] = r.Int64("id")
		mm["guid"] = r.String("guid")
		mm["username"] = r.String("username")
		mm["device_guid"] = r.String("device_guid")
		mm["device_name"] = r.String("device_name")
		mm["device_model"] = r.String("device_model")
		return mm
	}
}

func userWithDeviceSelect() (sq.Fields, func(context.Context, *sq.Row) *UserWithDevice) {
	tbl := UserTable
	joinTbl := UserDeviceTable
	fields := sq.Fields{
		tbl.ID.As("user_id"),
		tbl.GUID.As("user_guid"),
		tbl.USERNAME,
		joinTbl.GUID.As("device_guid"),
		joinTbl.NAME.As("device_name"),
		joinTbl.MODEL.As("device_model"),
	}
	return fields, userJoinDeviceRowMapper()
}

func userJoinDeviceRowMapper() func(context.Context, *sq.Row) *UserWithDevice {
	return func(ctx context.Context, r *sq.Row) *UserWithDevice {
		v := &UserWithDevice{
			UserID:      r.Int64("user_id"),
			UserGUID:    r.String("user_guid"),
			Username:    r.String("username"),
			DeviceGUID:  r.String("device_guid"),
			DeviceName:  r.String("device_name"),
			DeviceModel: r.String("device_model"),
		}
		return v
	}
}

func userModelRowMapper() func(context.Context, *sq.Row) *User {
	return func(ctx context.Context, r *sq.Row) *User {
		tbl := Tables.Users

		u := &User{
			ID:        r.Int64Field(tbl.ID),
			GUID:      r.StringField(tbl.GUID),
			Username:  r.StringField(tbl.USERNAME),
			FirstName: r.StringField(tbl.FIRST_NAME),
			LastName:  r.StringField(tbl.LAST_NAME),
			Score:     r.Float64Field(tbl.SCORE),
			Level:     r.Int16Field(tbl.LEVEL),
			TenantID:  r.Int64Field(tbl.TENANT_ID),
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

func userInsertColumnMapper(col *sq.Column, r User) {
	tbl := UserTable

	if r.ID > 0 {
		col.SetInt64(tbl.ID, r.ID)
	}
	col.Set(tbl.GUID, r.GUID)
	col.Set(tbl.USERNAME, r.Username)
	col.Set(tbl.SCORE, r.Score)
	col.Set(tbl.CREATED_AT, r.CreatedAt)
	col.Set(tbl.UPDATED_AT, r.UpdatedAt)
}

func userDeviceInsertColumnMapper(col *sq.Column, r *UserDevice) {
	tbl := UserDeviceTable

	col.Set(tbl.USER_ID, r.UserID)
	col.Set(tbl.GUID, r.GUID)
	col.Set(tbl.NAME, r.Name)
	col.Set(tbl.MODEL, r.Model)
	col.SetTime(tbl.CREATED_AT, r.CreatedAt)
	col.SetTime(tbl.UPDATED_AT, r.UpdatedAt)
}

func deviceInsertColumnMapper(col *sq.Column, r *Device) {
	tbl := DeviceTable

	col.SetString(tbl.GUID, r.GUID)
	col.SetString(tbl.NAME, r.Name)
	col.SetTime(tbl.CREATED_AT, r.CreatedAt)
	col.SetTime(tbl.UPDATED_AT, r.UpdatedAt)
}

func deviceRowMapper() func(context.Context, *sq.Row) *Device {
	return func(ctx context.Context, r *sq.Row) *Device {
		tbl := DeviceTable

		u := &Device{
			ID:        r.Int64Field(tbl.ID),
			GUID:      r.StringField(tbl.GUID),
			CreatedAt: r.TimeField(tbl.CREATED_AT),
			UpdatedAt: r.TimeField(tbl.UPDATED_AT),
		}

		return u
	}
}

func randomUser() User {
	ln := time.Now().Local()
	u := User{
		GUID:      gofakeit.UUID(),
		Username:  gofakeit.Username(),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Score:     gofakeit.Float64(),
		Level:     int16(gofakeit.IntRange(1, 99)),
		CreatedAt: ln,
		UpdatedAt: ln,
		TenantID:  int64(gofakeit.IntRange(1, 5)),
	}
	return u
}

func randomUserSetter() UserSetter {
	ln := time.Now().Local()
	u := UserSetter{
		GUID:      omit.From(gofakeit.UUID()),
		Username:  omit.From(gofakeit.Username()),
		FirstName: omit.From(gofakeit.FirstName()),
		LastName:  omit.From(gofakeit.LastName()),
		Score:     omit.From(gofakeit.Float64()),
		Level:     omit.From(int16(gofakeit.IntRange(1, 99))),
		CreatedAt: omit.From(ln),
		UpdatedAt: omit.From(ln),
		TenantID:  omit.From(int64(gofakeit.IntRange(1, 5))),
	}
	return u
}

func randomUserDevice() *UserDevice {
	ln := time.Now().Local()
	u := &UserDevice{
		UserID:    int64(gofakeit.IntRange(1, 30)),
		GUID:      gofakeit.UUID(),
		Name:      gofakeit.AppName(),
		Model:     gofakeit.CarModel(),
		CreatedAt: ln,
		UpdatedAt: ln,
	}
	return u
}

func randomDevice() *Device {
	ln := time.Now().Local()
	u := &Device{
		GUID:      gofakeit.UUID(),
		Name:      gofakeit.AppName(),
		CreatedAt: ln,
		UpdatedAt: ln,
	}
	return u
}

func randomTag(desc *string) Tag {
	u := Tag{
		GUID:        gofakeit.UUID(),
		Code:        gofakeit.City(),
		Name:        gofakeit.DomainName(),
		Description: null.FromPtr(desc),
		CreatedAt:   time.Now().Local(),
	}
	return u
}

func randomTagSetter(desc *string) TagSetter {
	u := TagSetter{
		GUID:        omit.From(gofakeit.UUID()),
		Code:        omit.From(gofakeit.City()),
		Name:        omit.From(gofakeit.DomainName()),
		Description: omitnull.FromPtr(desc),
		CreatedAt:   omit.From(time.Now().Local()),
	}
	return u
}

type Mod[T any] interface {
	Apply(T)
}
