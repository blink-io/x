package orm

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/blink-io/opt/null"
	"github.com/blink-io/opt/omit"
	"github.com/blink-io/opt/omitnull"
	"github.com/blink-io/x/log/slog/handlers/color"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sanity-io/litter"
)

var _log = slog.New(color.New(os.Stderr, color.Options{
	Level: slog.LevelDebug,
}))

var ctx = context.Background()

func init() {
	l := sq.NewLogger(os.Stdout, "", log.LstdFlags, sq.LoggerConfig{
		ShowCaller:    true,
		ShowTimeTaken: true,
		HideArgs:      true,
	})
	sq.SetDefaultLogQuery(func(ctx context.Context, stats sq.QueryStats) {
		l.SqLogQuery(ctx, stats)
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
	UserStatusActive  UserStatus = "active"
	UserStatusBlocked UserStatus = "blocked"
)

func (v UserStatus) String() string {
	return string(v)
}

func (m User) String() string {
	return litter.Sdump(m)
}

type UserWithDevice struct {
	ID          int64  `db:"id"`
	GUID        string `db:"guid"`
	Username    string `db:"username"`
	DeviceGUID  string `db:"device_guid"`
	DeviceName  string `db:"device_name"`
	DeviceModel string `db:"device_model"`
}

func (m UserWithDevice) String() string {
	return litter.Sdump(m)
}

type UserSetter struct {
	ID        omit.Val[int]       `db:"id"`
	GUID      omit.Val[string]    `db:"guid"`
	Username  omit.Val[string]    `db:"username"`
	Score     omit.Val[float64]   `db:"score"`
	CreatedAt omit.Val[time.Time] `db:"created_at"`
	UpdatedAt omit.Val[time.Time] `db:"updated_at"`
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
				Where(tbl.ID.EqInt(id)),
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
		c.SetInt(tbl.ID, v)
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

func (m UserSetter) SetColumns(c *sq.Column) {
	m.setColumns(c, true)
}

type UserDevice struct {
	ID          int              `db:"id"`
	GUID        string           `db:"guid"`
	UserID      int              `db:"user_id"`
	Name        string           `db:"name"`
	Model       string           `db:"score"`
	Description null.Val[string] `db:"description"`
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
}

func (m UserDevice) String() string {
	return litter.Sdump(m)
}

type Device struct {
	ID        int       `db:"id"`
	GUID      string    `db:"guid"`
	Name      string    `db:"name"`
	Model     string    `db:"model"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m Device) Insert(db sq.DB) error {
	tbl := DeviceTable
	_, err := sq.Exec(db, sq.InsertInto(tbl).ColumnValues(func(c *sq.Column) {
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

type Tag struct {
	ID          int              `db:"id"`
	GUID        string           `db:"guid"`
	Code        string           `db:"code"`
	Name        string           `db:"name"`
	CreatedAt   time.Time        `db:"created_at"`
	Description null.Val[string] `db:"description"`
}

func (m Tag) Setter() TagSetter {
	ss := TagSetter{
		GUID:        omit.From(m.GUID),
		Code:        omit.From(m.Code),
		Name:        omit.From(m.Name),
		CreatedAt:   omit.From(m.CreatedAt),
		Description: omitnull.FromNull(m.Description),
	}
	if m.ID > 0 {
		ss.ID = omit.From(m.ID)
	}
	return ss
}

type TagSetter struct {
	ID          omit.Val[int]        `db:"id"`
	GUID        omit.Val[string]     `db:"guid"`
	Code        omit.Val[string]     `db:"code"`
	Name        omit.Val[string]     `db:"name"`
	CreatedAt   omit.Val[time.Time]  `db:"created_at"`
	Description omitnull.Val[string] `db:"description"`
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
		//tbl := UserTable
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

func userWithDeviceSelect() (sq.Fields, func(*sq.Row) *UserWithDevice) {
	tbl := UserTable
	joinTbl := UserDeviceTable
	fields := sq.Fields{
		tbl.ID,
		tbl.GUID,
		tbl.USERNAME,
		joinTbl.GUID.As("device_guid"),
		joinTbl.NAME.As("device_name"),
		joinTbl.MODEL.As("device_model"),
	}
	return fields, userJoinDeviceRowMapper()
}

func userJoinDeviceRowMapper() func(*sq.Row) *UserWithDevice {
	return func(r *sq.Row) *UserWithDevice {
		v := &UserWithDevice{
			ID:          r.Int64("id"),
			GUID:        r.String("guid"),
			Username:    r.String("username"),
			DeviceGUID:  r.String("device_guid"),
			DeviceName:  r.String("device_name"),
			DeviceModel: r.String("device_model"),
		}
		return v
	}
}

func userModelRowMapper() func(*sq.Row) *User {
	return func(r *sq.Row) *User {
		tbl := Tables.Users

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

func userInsertColumnMapper(col *sq.Column, r User) {
	tbl := UserTable

	if r.ID > 0 {
		col.SetInt(tbl.ID, r.ID)
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

func deviceRowMapper() func(*sq.Row) *Device {
	return func(r *sq.Row) *Device {
		tbl := DeviceTable

		u := &Device{
			ID:        r.IntField(tbl.ID),
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
		Level:     gofakeit.IntRange(1, 99),
		CreatedAt: ln,
		UpdatedAt: ln,
		TenantID:  gofakeit.IntRange(1, 5),
	}
	return u
}

func randomUserDevice() *UserDevice {
	ln := time.Now().Local()
	u := &UserDevice{
		UserID:    gofakeit.IntRange(1, 30),
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

type Mod[T any] interface {
	Apply(T)
}
