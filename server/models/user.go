// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// User is an object representing the database table.
type User struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Icon      string `boil:"icon" json:"icon" toml:"icon" yaml:"icon"`
	IsAdmin   string `boil:"is_admin" json:"is_admin" toml:"is_admin" yaml:"is_admin"`
	CreatedAt string `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt string `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserColumns = struct {
	ID        string
	Name      string
	Icon      string
	IsAdmin   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Icon:      "icon",
	IsAdmin:   "is_admin",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserTableColumns = struct {
	ID        string
	Name      string
	Icon      string
	IsAdmin   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "user.id",
	Name:      "user.name",
	Icon:      "user.icon",
	IsAdmin:   "user.is_admin",
	CreatedAt: "user.created_at",
	UpdatedAt: "user.updated_at",
}

// Generated where

var UserWhere = struct {
	ID        whereHelperstring
	Name      whereHelperstring
	Icon      whereHelperstring
	IsAdmin   whereHelperstring
	CreatedAt whereHelperstring
	UpdatedAt whereHelperstring
}{
	ID:        whereHelperstring{field: "\"user\".\"id\""},
	Name:      whereHelperstring{field: "\"user\".\"name\""},
	Icon:      whereHelperstring{field: "\"user\".\"icon\""},
	IsAdmin:   whereHelperstring{field: "\"user\".\"is_admin\""},
	CreatedAt: whereHelperstring{field: "\"user\".\"created_at\""},
	UpdatedAt: whereHelperstring{field: "\"user\".\"updated_at\""},
}

// UserRels is where relationship names are stored.
var UserRels = struct {
	OwnerContests string
	Entries       string
}{
	OwnerContests: "OwnerContests",
	Entries:       "Entries",
}

// userR is where relationships are stored.
type userR struct {
	OwnerContests ContestSlice `boil:"OwnerContests" json:"OwnerContests" toml:"OwnerContests" yaml:"OwnerContests"`
	Entries       EntrySlice   `boil:"Entries" json:"Entries" toml:"Entries" yaml:"Entries"`
}

// NewStruct creates a new relationship struct
func (*userR) NewStruct() *userR {
	return &userR{}
}

func (r *userR) GetOwnerContests() ContestSlice {
	if r == nil {
		return nil
	}
	return r.OwnerContests
}

func (r *userR) GetEntries() EntrySlice {
	if r == nil {
		return nil
	}
	return r.Entries
}

// userL is where Load methods for each relationship are stored.
type userL struct{}

var (
	userAllColumns            = []string{"id", "name", "icon", "is_admin", "created_at", "updated_at"}
	userColumnsWithoutDefault = []string{"id", "name", "icon", "created_at", "updated_at"}
	userColumnsWithDefault    = []string{"is_admin"}
	userPrimaryKeyColumns     = []string{"id"}
	userGeneratedColumns      = []string{}
)

type (
	// UserSlice is an alias for a slice of pointers to User.
	// This should almost always be used instead of []User.
	UserSlice []*User
	// UserHook is the signature for custom User hook methods
	UserHook func(context.Context, boil.ContextExecutor, *User) error

	userQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userType                 = reflect.TypeOf(&User{})
	userMapping              = queries.MakeStructMapping(userType)
	userPrimaryKeyMapping, _ = queries.BindMapping(userType, userMapping, userPrimaryKeyColumns)
	userInsertCacheMut       sync.RWMutex
	userInsertCache          = make(map[string]insertCache)
	userUpdateCacheMut       sync.RWMutex
	userUpdateCache          = make(map[string]updateCache)
	userUpsertCacheMut       sync.RWMutex
	userUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userAfterSelectHooks []UserHook

var userBeforeInsertHooks []UserHook
var userAfterInsertHooks []UserHook

var userBeforeUpdateHooks []UserHook
var userAfterUpdateHooks []UserHook

var userBeforeDeleteHooks []UserHook
var userAfterDeleteHooks []UserHook

var userBeforeUpsertHooks []UserHook
var userAfterUpsertHooks []UserHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *User) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *User) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *User) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *User) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *User) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *User) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *User) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *User) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *User) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserHook registers your hook function for all future operations.
func AddUserHook(hookPoint boil.HookPoint, userHook UserHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userAfterSelectHooks = append(userAfterSelectHooks, userHook)
	case boil.BeforeInsertHook:
		userBeforeInsertHooks = append(userBeforeInsertHooks, userHook)
	case boil.AfterInsertHook:
		userAfterInsertHooks = append(userAfterInsertHooks, userHook)
	case boil.BeforeUpdateHook:
		userBeforeUpdateHooks = append(userBeforeUpdateHooks, userHook)
	case boil.AfterUpdateHook:
		userAfterUpdateHooks = append(userAfterUpdateHooks, userHook)
	case boil.BeforeDeleteHook:
		userBeforeDeleteHooks = append(userBeforeDeleteHooks, userHook)
	case boil.AfterDeleteHook:
		userAfterDeleteHooks = append(userAfterDeleteHooks, userHook)
	case boil.BeforeUpsertHook:
		userBeforeUpsertHooks = append(userBeforeUpsertHooks, userHook)
	case boil.AfterUpsertHook:
		userAfterUpsertHooks = append(userAfterUpsertHooks, userHook)
	}
}

// OneG returns a single user record from the query using the global executor.
func (q userQuery) OneG(ctx context.Context) (*User, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single user record from the query.
func (q userQuery) One(ctx context.Context, exec boil.ContextExecutor) (*User, error) {
	o := &User{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all User records from the query using the global executor.
func (q userQuery) AllG(ctx context.Context) (UserSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all User records from the query.
func (q userQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserSlice, error) {
	var o []*User

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to User slice")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all User records in the query using the global executor
func (q userQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all User records in the query.
func (q userQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q userQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q userQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user exists")
	}

	return count > 0, nil
}

// OwnerContests retrieves all the contest's Contests with an executor via owner column.
func (o *User) OwnerContests(mods ...qm.QueryMod) contestQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"contest\".\"owner\"=?", o.ID),
	)

	return Contests(queryMods...)
}

// Entries retrieves all the entry's Entries with an executor.
func (o *User) Entries(mods ...qm.QueryMod) entryQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"entry\".\"user_id\"=?", o.ID),
	)

	return Entries(queryMods...)
}

// LoadOwnerContests allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userL) LoadOwnerContests(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUser interface{}, mods queries.Applicator) error {
	var slice []*User
	var object *User

	if singular {
		var ok bool
		object, ok = maybeUser.(*User)
		if !ok {
			object = new(User)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUser))
			}
		}
	} else {
		s, ok := maybeUser.(*[]*User)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUser))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`contest`),
		qm.WhereIn(`contest.owner in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load contest")
	}

	var resultSlice []*Contest
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice contest")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on contest")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for contest")
	}

	if len(contestAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.OwnerContests = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &contestR{}
			}
			foreign.R.OwnerUser = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.Owner {
				local.R.OwnerContests = append(local.R.OwnerContests, foreign)
				if foreign.R == nil {
					foreign.R = &contestR{}
				}
				foreign.R.OwnerUser = local
				break
			}
		}
	}

	return nil
}

// LoadEntries allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (userL) LoadEntries(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUser interface{}, mods queries.Applicator) error {
	var slice []*User
	var object *User

	if singular {
		var ok bool
		object, ok = maybeUser.(*User)
		if !ok {
			object = new(User)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUser))
			}
		}
	} else {
		s, ok := maybeUser.(*[]*User)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUser))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`entry`),
		qm.WhereIn(`entry.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load entry")
	}

	var resultSlice []*Entry
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice entry")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on entry")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for entry")
	}

	if len(entryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Entries = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &entryR{}
			}
			foreign.R.User = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.UserID {
				local.R.Entries = append(local.R.Entries, foreign)
				if foreign.R == nil {
					foreign.R = &entryR{}
				}
				foreign.R.User = local
				break
			}
		}
	}

	return nil
}

// AddOwnerContestsG adds the given related objects to the existing relationships
// of the user, optionally inserting them as new records.
// Appends related to o.R.OwnerContests.
// Sets related.R.OwnerUser appropriately.
// Uses the global database handle.
func (o *User) AddOwnerContestsG(ctx context.Context, insert bool, related ...*Contest) error {
	return o.AddOwnerContests(ctx, boil.GetContextDB(), insert, related...)
}

// AddOwnerContests adds the given related objects to the existing relationships
// of the user, optionally inserting them as new records.
// Appends related to o.R.OwnerContests.
// Sets related.R.OwnerUser appropriately.
func (o *User) AddOwnerContests(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Contest) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.Owner = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"contest\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"owner"}),
				strmangle.WhereClause("\"", "\"", 0, contestPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.Owner = o.ID
		}
	}

	if o.R == nil {
		o.R = &userR{
			OwnerContests: related,
		}
	} else {
		o.R.OwnerContests = append(o.R.OwnerContests, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contestR{
				OwnerUser: o,
			}
		} else {
			rel.R.OwnerUser = o
		}
	}
	return nil
}

// AddEntriesG adds the given related objects to the existing relationships
// of the user, optionally inserting them as new records.
// Appends related to o.R.Entries.
// Sets related.R.User appropriately.
// Uses the global database handle.
func (o *User) AddEntriesG(ctx context.Context, insert bool, related ...*Entry) error {
	return o.AddEntries(ctx, boil.GetContextDB(), insert, related...)
}

// AddEntries adds the given related objects to the existing relationships
// of the user, optionally inserting them as new records.
// Appends related to o.R.Entries.
// Sets related.R.User appropriately.
func (o *User) AddEntries(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Entry) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.UserID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"entry\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"user_id"}),
				strmangle.WhereClause("\"", "\"", 0, entryPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.UserID = o.ID
		}
	}

	if o.R == nil {
		o.R = &userR{
			Entries: related,
		}
	} else {
		o.R.Entries = append(o.R.Entries, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &entryR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// Users retrieves all the records using an executor.
func Users(mods ...qm.QueryMod) userQuery {
	mods = append(mods, qm.From("\"user\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user\".*"})
	}

	return userQuery{q}
}

// FindUserG retrieves a single record by ID.
func FindUserG(ctx context.Context, iD string, selectCols ...string) (*User, error) {
	return FindUser(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUser(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*User, error) {
	userObj := &User{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user")
	}

	if err = userObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userObj, err
	}

	return userObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *User) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *User) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userInsertCacheMut.RLock()
	cache, cached := userInsertCache[key]
	userInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userAllColumns,
			userColumnsWithDefault,
			userColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userType, userMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userType, userMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user")
	}

	if !cached {
		userInsertCacheMut.Lock()
		userInsertCache[key] = cache
		userInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single User record using the global executor.
// See Update for more documentation.
func (o *User) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the User.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *User) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userUpdateCacheMut.RLock()
	cache, cached := userUpdateCache[key]
	userUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userAllColumns,
			userPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, userPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userType, userMapping, append(wl, userPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update user row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user")
	}

	if !cached {
		userUpdateCacheMut.Lock()
		userUpdateCache[key] = cache
		userUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q userQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in user slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all user")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *User) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *User) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userUpsertCacheMut.RLock()
	cache, cached := userUpsertCache[key]
	userUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userAllColumns,
			userColumnsWithDefault,
			userColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userAllColumns,
			userPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert user, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userPrimaryKeyColumns))
			copy(conflict, userPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"user\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userType, userMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userType, userMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert user")
	}

	if !cached {
		userUpsertCacheMut.Lock()
		userUpsertCache[key] = cache
		userUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single User record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *User) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single User record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *User) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no User provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userPrimaryKeyMapping)
	sql := "DELETE FROM \"user\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q userQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q userQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user")
	}

	if len(userAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *User) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no User provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *User) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty UserSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user\".* FROM \"user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserSlice")
	}

	*o = slice

	return nil
}

// UserExistsG checks if the User row exists.
func UserExistsG(ctx context.Context, iD string) (bool, error) {
	return UserExists(ctx, boil.GetContextDB(), iD)
}

// UserExists checks if the User row exists.
func UserExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user exists")
	}

	return exists, nil
}
