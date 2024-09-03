package orm

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sanity-io/litter"
)

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

type User struct {
	ID        int       `db:"id"`
	GUID      string    `db:"guid"`
	Username  string    `db:"username"`
	Score     float64   `db:"score"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
	ID          int                  `db:"id"`
	GUID        string               `db:"guid"`
	UserID      int                  `db:"user_id"`
	Name        string               `db:"name"`
	Model       string               `db:"score"`
	Description omitnull.Val[string] `db:"description"`
	CreatedAt   time.Time            `db:"created_at"`
	UpdatedAt   time.Time            `db:"updated_at"`
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
	ID          int                  `db:"id"`
	GUID        string               `db:"guid"`
	Code        string               `db:"code"`
	Name        string               `db:"name"`
	Description omitnull.Val[string] `db:"description"`
}

func (m Tag) Insert(db sq.DB) error {
	tbl := TagTable
	_, err := sq.Exec(db, sq.InsertInto(tbl).ColumnValues(func(c *sq.Column) {
		c.SetString(tbl.GUID, m.GUID)
		c.SetString(tbl.NAME, m.Name)
		c.SetString(tbl.CODE, m.Code)
		if v, ok := m.Description.Get(); ok {
			c.SetString(tbl.DESCRIPTION, v)
		}
	}))
	return err
}

type Model struct {
	Name    string `db:"name"`
	Version string `db:"version"`
	Current string `db:"current"`
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
		tbl := UserTable

		u := &User{
			ID:       r.IntField(tbl.ID),
			GUID:     r.StringField(tbl.GUID),
			Username: r.StringField(tbl.USERNAME),
			Score:    r.Float64Field(tbl.SCORE),
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

		panic(errors.New("my custom error"))

		return u
	}
}

func userInsertColumnMapper(col *sq.Column, r *User) {
	tbl := UserTable

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

func tagInsertColumnMapper(col *sq.Column, r *Tag) {
	tbl := TagTable

	col.SetString(tbl.GUID, r.GUID)
	col.SetString(tbl.NAME, r.Name)
	col.SetString(tbl.CODE, r.Code)
	col.Set(tbl.DESCRIPTION, r.Description)
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

func randomUser() *User {
	ln := time.Now().Local()
	u := &User{
		GUID:      gofakeit.UUID(),
		Username:  gofakeit.Username(),
		Score:     gofakeit.Float64(),
		CreatedAt: ln,
		UpdatedAt: ln,
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

func randomTag(desc *string) *Tag {
	u := &Tag{
		GUID:        gofakeit.UUID(),
		Code:        gofakeit.City(),
		Name:        gofakeit.DomainName(),
		Description: omitnull.FromPtr(desc),
	}
	return u
}
