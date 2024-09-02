package orm

import (
	"errors"
	"fmt"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/bokwoon95/sq"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sanity-io/litter"
)

type UserTable struct {
	sq.TableStruct `sq:"users"`
	ID             sq.NumberField `sq:"id"`
	GUID           sq.StringField `sq:"guid"`
	Username       sq.StringField `sq:"username"`
	Score          sq.NumberField `sq:"score"`
	CreatedAt      sq.TimeField   `sq:"created_at"`
	UpdatedAt      sq.TimeField   `sq:"updated_at"`
}

type UserDeviceTable struct {
	sq.TableStruct `sq:"user_devices"`
	ID             sq.NumberField `sq:"id"`
	GUID           sq.StringField `sq:"guid"`
	UserID         sq.NumberField `sq:"user_id"`
	Device         sq.StringField `sq:"device"`
	Model          sq.StringField `sq:"model"`
	CreatedAt      sq.TimeField   `sq:"created_at"`
	UpdatedAt      sq.TimeField   `sq:"updated_at"`
}

type DeviceTable struct {
	sq.TableStruct `sq:"devices"`
	ID             sq.NumberField `sq:"id"`
	GUID           sq.StringField `sq:"guid"`
	Name           sq.StringField `sq:"name"`
	CreatedAtTS    sq.TimeField   `sq:"created_at_ts"`
	UpdatedAtTS    sq.TimeField   `sq:"updated_at_ts"`
}

type TagTable struct {
	sq.TableStruct `sq:"tags"`
	ID             sq.NumberField `sq:"id"`
	GUID           sq.StringField `sq:"guid"`
	Code           sq.StringField `sq:"code"`
	Name           sq.StringField `sq:"name"`
	Description    sq.StringField `sq:"description"`
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

func (m UserSetter) Insert(db sq.DB) error {
	tbl := UserTableDef
	_, err := sq.Exec(
		db,
		sq.InsertInto(tbl).
			ColumnValues(func(c *sq.Column) {
				m.setColumns(c, false)
			}),
	)
	return err
}

func (m UserSetter) UpdateByID(db sq.DB) error {
	tbl := UserTableDef
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

func (m UserSetter) UpdateByWhere(db sq.DB, where ...sq.Predicate) error {
	if len(where) == 0 {
		return errors.New("where is empty")
	}
	tbl := UserTableDef
	_, err := sq.Exec(
		db,
		sq.Update(tbl).
			SetFunc(m.SetColumns).
			Where(where...),
	)
	return err

}

func (m UserSetter) setColumns(c *sq.Column, withID bool) {
	tbl := UserTableDef
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
		c.SetString(tbl.Username, v)
	}
	if !m.Score.IsUnset() {
		v, _ := m.Score.Get()
		c.SetFloat64(tbl.Score, v)
	}
	if !m.CreatedAt.IsUnset() {
		v, _ := m.CreatedAt.Get()
		c.SetTime(tbl.CreatedAt, v)
	}
	if !m.UpdatedAt.IsUnset() {
		v, _ := m.UpdatedAt.Get()
		c.SetTime(tbl.UpdatedAt, v)
	}
}

func (m UserSetter) SetColumns(c *sq.Column) {
	m.setColumns(c, true)
}

type UserDevice struct {
	ID        int       `db:"id"`
	GUID      string    `db:"guid"`
	UserID    int       `db:"user_id"`
	Device    string    `db:"device"`
	Model     string    `db:"score"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m UserDevice) String() string {
	return litter.Sdump(m)
}

type Device struct {
	ID          int       `db:"id"`
	GUID        string    `db:"guid"`
	Name        string    `db:"name"`
	CreatedAtTS time.Time `db:"created_at_ts"`
	UpdatedAtTS time.Time `db:"updated_at_ts"`
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

type Model struct {
	Name    string `db:"name"`
	Version string `db:"version"`
	Current string `db:"current"`
}

func (m Model) String() string {
	return litter.Sdump(m)
}

var _ fmt.Stringer = (*User)(nil)
var _ fmt.Stringer = (*UserDevice)(nil)

var (
	UserTableDef       = sq.New[UserTable]("u1")
	UserDeviceTableDef = sq.New[UserDeviceTable]("u2")
	DeviceTableDef     = sq.New[DeviceTable]("u3")
	TagTableDef        = sq.New[TagTable]("u4")
)

func userModelRowMapper() func(*sq.Row) *User {
	return func(r *sq.Row) *User {
		tbl := UserTableDef

		u := &User{
			ID:       r.IntField(tbl.ID),
			GUID:     r.StringField(tbl.GUID),
			Username: r.StringField(tbl.Username),
			Score:    r.Float64Field(tbl.Score),
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
			u.CreatedAt = r.TimeField(tbl.CreatedAt)
			u.UpdatedAt = r.TimeField(tbl.UpdatedAt)
		}

		return u
	}
}

func userInsertColumnMapper(col *sq.Column, r *User) {
	tbl := UserTableDef

	col.Set(tbl.GUID, r.GUID)
	col.Set(tbl.Username, r.Username)
	col.Set(tbl.Score, r.Score)
	col.Set(tbl.CreatedAt, r.CreatedAt)
	col.Set(tbl.UpdatedAt, r.UpdatedAt)
}

func userDeviceInsertColumnMapper(col *sq.Column, r *UserDevice) {
	tbl := UserDeviceTableDef

	col.Set(tbl.UserID, r.UserID)
	col.Set(tbl.GUID, r.GUID)
	col.Set(tbl.Device, r.Device)
	col.Set(tbl.Model, r.Model)
	col.SetTime(tbl.CreatedAt, r.CreatedAt)
	col.SetTime(tbl.UpdatedAt, r.UpdatedAt)
}

func deviceInsertColumnMapper(col *sq.Column, r *Device) {
	tbl := DeviceTableDef

	col.SetString(tbl.GUID, r.GUID)
	col.SetString(tbl.Name, r.Name)
	col.SetTime(tbl.CreatedAtTS, r.CreatedAtTS)
	col.SetTime(tbl.UpdatedAtTS, r.UpdatedAtTS)
}

func tagInsertColumnMapper(col *sq.Column, r *Tag) {
	tbl := TagTableDef

	col.SetString(tbl.GUID, r.GUID)
	col.SetString(tbl.Name, r.Name)
	col.SetString(tbl.Code, r.Code)
	col.Set(tbl.Description, r.Description)
}

func deviceRowMapper() func(*sq.Row) *Device {
	return func(r *sq.Row) *Device {
		tbl := DeviceTableDef

		cc := r.Int64("created_at_ts")
		uu := r.Int64("updated_at_ts")

		aa := time.UnixMilli(cc)
		bb := time.UnixMilli(uu)
		u := &Device{
			ID:          r.IntField(tbl.ID),
			GUID:        r.StringField(tbl.GUID),
			CreatedAtTS: aa,
			UpdatedAtTS: bb,
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
		Device:    gofakeit.AppName(),
		Model:     gofakeit.CarModel(),
		CreatedAt: ln,
		UpdatedAt: ln,
	}
	return u
}

func randomDevice() *Device {
	ln := time.Now().Local()
	u := &Device{
		GUID:        gofakeit.UUID(),
		Name:        gofakeit.AppName(),
		CreatedAtTS: ln,
		UpdatedAtTS: ln,
	}
	return u
}

func randomTag(desc *string) *Tag {
	u := &Tag{
		GUID:        gofakeit.UUID(),
		Code:        gofakeit.Fruit(),
		Name:        gofakeit.DomainName(),
		Description: omitnull.FromPtr(desc),
	}
	return u
}
