// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)

// WithdrawalFiat is an object representing the database table.
type WithdrawalFiat struct {
	ID                  int64   `boil:"id" json:"id" toml:"id" yaml:"id"`
	BankName            string  `boil:"bank_name" json:"bank_name" toml:"bank_name" yaml:"bank_name"`
	BankAddress         string  `boil:"bank_address" json:"bank_address" toml:"bank_address" yaml:"bank_address"`
	BankAccountName     string  `boil:"bank_account_name" json:"bank_account_name" toml:"bank_account_name" yaml:"bank_account_name"`
	BankAccountNumber   string  `boil:"bank_account_number" json:"bank_account_number" toml:"bank_account_number" yaml:"bank_account_number"`
	BSB                 string  `boil:"bsb" json:"bsb" toml:"bsb" yaml:"bsb"`
	SwiftCode           string  `boil:"swift_code" json:"swift_code" toml:"swift_code" yaml:"swift_code"`
	Iban                string  `boil:"iban" json:"iban" toml:"iban" yaml:"iban"`
	BankCode            float64 `boil:"bank_code" json:"bank_code" toml:"bank_code" yaml:"bank_code"`
	WithdrawalHistoryID string  `boil:"withdrawal_history_id" json:"withdrawal_history_id" toml:"withdrawal_history_id" yaml:"withdrawal_history_id"`

	R *withdrawalFiatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L withdrawalFiatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WithdrawalFiatColumns = struct {
	ID                  string
	BankName            string
	BankAddress         string
	BankAccountName     string
	BankAccountNumber   string
	BSB                 string
	SwiftCode           string
	Iban                string
	BankCode            string
	WithdrawalHistoryID string
}{
	ID:                  "id",
	BankName:            "bank_name",
	BankAddress:         "bank_address",
	BankAccountName:     "bank_account_name",
	BankAccountNumber:   "bank_account_number",
	BSB:                 "bsb",
	SwiftCode:           "swift_code",
	Iban:                "iban",
	BankCode:            "bank_code",
	WithdrawalHistoryID: "withdrawal_history_id",
}

// Generated where

var WithdrawalFiatWhere = struct {
	ID                  whereHelperint64
	BankName            whereHelperstring
	BankAddress         whereHelperstring
	BankAccountName     whereHelperstring
	BankAccountNumber   whereHelperstring
	BSB                 whereHelperstring
	SwiftCode           whereHelperstring
	Iban                whereHelperstring
	BankCode            whereHelperfloat64
	WithdrawalHistoryID whereHelperstring
}{
	ID:                  whereHelperint64{field: "\"withdrawal_fiat\".\"id\""},
	BankName:            whereHelperstring{field: "\"withdrawal_fiat\".\"bank_name\""},
	BankAddress:         whereHelperstring{field: "\"withdrawal_fiat\".\"bank_address\""},
	BankAccountName:     whereHelperstring{field: "\"withdrawal_fiat\".\"bank_account_name\""},
	BankAccountNumber:   whereHelperstring{field: "\"withdrawal_fiat\".\"bank_account_number\""},
	BSB:                 whereHelperstring{field: "\"withdrawal_fiat\".\"bsb\""},
	SwiftCode:           whereHelperstring{field: "\"withdrawal_fiat\".\"swift_code\""},
	Iban:                whereHelperstring{field: "\"withdrawal_fiat\".\"iban\""},
	BankCode:            whereHelperfloat64{field: "\"withdrawal_fiat\".\"bank_code\""},
	WithdrawalHistoryID: whereHelperstring{field: "\"withdrawal_fiat\".\"withdrawal_history_id\""},
}

// WithdrawalFiatRels is where relationship names are stored.
var WithdrawalFiatRels = struct {
	WithdrawalHistory string
}{
	WithdrawalHistory: "WithdrawalHistory",
}

// withdrawalFiatR is where relationships are stored.
type withdrawalFiatR struct {
	WithdrawalHistory *WithdrawalHistory
}

// NewStruct creates a new relationship struct
func (*withdrawalFiatR) NewStruct() *withdrawalFiatR {
	return &withdrawalFiatR{}
}

// withdrawalFiatL is where Load methods for each relationship are stored.
type withdrawalFiatL struct{}

var (
	withdrawalFiatAllColumns            = []string{"id", "bank_name", "bank_address", "bank_account_name", "bank_account_number", "bsb", "swift_code", "iban", "bank_code", "withdrawal_history_id"}
	withdrawalFiatColumnsWithoutDefault = []string{"bank_name", "bank_address", "bank_account_name", "bank_account_number", "bsb", "swift_code", "iban", "bank_code", "withdrawal_history_id"}
	withdrawalFiatColumnsWithDefault    = []string{"id"}
	withdrawalFiatPrimaryKeyColumns     = []string{"id"}
)

type (
	// WithdrawalFiatSlice is an alias for a slice of pointers to WithdrawalFiat.
	// This should generally be used opposed to []WithdrawalFiat.
	WithdrawalFiatSlice []*WithdrawalFiat
	// WithdrawalFiatHook is the signature for custom WithdrawalFiat hook methods
	WithdrawalFiatHook func(context.Context, boil.ContextExecutor, *WithdrawalFiat) error

	withdrawalFiatQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	withdrawalFiatType                 = reflect.TypeOf(&WithdrawalFiat{})
	withdrawalFiatMapping              = queries.MakeStructMapping(withdrawalFiatType)
	withdrawalFiatPrimaryKeyMapping, _ = queries.BindMapping(withdrawalFiatType, withdrawalFiatMapping, withdrawalFiatPrimaryKeyColumns)
	withdrawalFiatInsertCacheMut       sync.RWMutex
	withdrawalFiatInsertCache          = make(map[string]insertCache)
	withdrawalFiatUpdateCacheMut       sync.RWMutex
	withdrawalFiatUpdateCache          = make(map[string]updateCache)
	withdrawalFiatUpsertCacheMut       sync.RWMutex
	withdrawalFiatUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var withdrawalFiatBeforeInsertHooks []WithdrawalFiatHook
var withdrawalFiatBeforeUpdateHooks []WithdrawalFiatHook
var withdrawalFiatBeforeDeleteHooks []WithdrawalFiatHook
var withdrawalFiatBeforeUpsertHooks []WithdrawalFiatHook

var withdrawalFiatAfterInsertHooks []WithdrawalFiatHook
var withdrawalFiatAfterSelectHooks []WithdrawalFiatHook
var withdrawalFiatAfterUpdateHooks []WithdrawalFiatHook
var withdrawalFiatAfterDeleteHooks []WithdrawalFiatHook
var withdrawalFiatAfterUpsertHooks []WithdrawalFiatHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *WithdrawalFiat) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *WithdrawalFiat) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *WithdrawalFiat) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *WithdrawalFiat) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *WithdrawalFiat) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *WithdrawalFiat) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *WithdrawalFiat) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *WithdrawalFiat) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *WithdrawalFiat) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range withdrawalFiatAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWithdrawalFiatHook registers your hook function for all future operations.
func AddWithdrawalFiatHook(hookPoint boil.HookPoint, withdrawalFiatHook WithdrawalFiatHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		withdrawalFiatBeforeInsertHooks = append(withdrawalFiatBeforeInsertHooks, withdrawalFiatHook)
	case boil.BeforeUpdateHook:
		withdrawalFiatBeforeUpdateHooks = append(withdrawalFiatBeforeUpdateHooks, withdrawalFiatHook)
	case boil.BeforeDeleteHook:
		withdrawalFiatBeforeDeleteHooks = append(withdrawalFiatBeforeDeleteHooks, withdrawalFiatHook)
	case boil.BeforeUpsertHook:
		withdrawalFiatBeforeUpsertHooks = append(withdrawalFiatBeforeUpsertHooks, withdrawalFiatHook)
	case boil.AfterInsertHook:
		withdrawalFiatAfterInsertHooks = append(withdrawalFiatAfterInsertHooks, withdrawalFiatHook)
	case boil.AfterSelectHook:
		withdrawalFiatAfterSelectHooks = append(withdrawalFiatAfterSelectHooks, withdrawalFiatHook)
	case boil.AfterUpdateHook:
		withdrawalFiatAfterUpdateHooks = append(withdrawalFiatAfterUpdateHooks, withdrawalFiatHook)
	case boil.AfterDeleteHook:
		withdrawalFiatAfterDeleteHooks = append(withdrawalFiatAfterDeleteHooks, withdrawalFiatHook)
	case boil.AfterUpsertHook:
		withdrawalFiatAfterUpsertHooks = append(withdrawalFiatAfterUpsertHooks, withdrawalFiatHook)
	}
}

// One returns a single withdrawalFiat record from the query.
func (q withdrawalFiatQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WithdrawalFiat, error) {
	o := &WithdrawalFiat{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: failed to execute a one query for withdrawal_fiat")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all WithdrawalFiat records from the query.
func (q withdrawalFiatQuery) All(ctx context.Context, exec boil.ContextExecutor) (WithdrawalFiatSlice, error) {
	var o []*WithdrawalFiat

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlite3: failed to assign all query results to WithdrawalFiat slice")
	}

	if len(withdrawalFiatAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all WithdrawalFiat records in the query.
func (q withdrawalFiatQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to count withdrawal_fiat rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q withdrawalFiatQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: failed to check if withdrawal_fiat exists")
	}

	return count > 0, nil
}

// WithdrawalHistory pointed to by the foreign key.
func (o *WithdrawalFiat) WithdrawalHistory(mods ...qm.QueryMod) withdrawalHistoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.WithdrawalHistoryID),
	}

	queryMods = append(queryMods, mods...)

	query := WithdrawalHistories(queryMods...)
	queries.SetFrom(query.Query, "\"withdrawal_history\"")

	return query
}

// LoadWithdrawalHistory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (withdrawalFiatL) LoadWithdrawalHistory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWithdrawalFiat interface{}, mods queries.Applicator) error {
	var slice []*WithdrawalFiat
	var object *WithdrawalFiat

	if singular {
		object = maybeWithdrawalFiat.(*WithdrawalFiat)
	} else {
		slice = *maybeWithdrawalFiat.(*[]*WithdrawalFiat)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &withdrawalFiatR{}
		}
		args = append(args, object.WithdrawalHistoryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &withdrawalFiatR{}
			}

			for _, a := range args {
				if a == obj.WithdrawalHistoryID {
					continue Outer
				}
			}

			args = append(args, obj.WithdrawalHistoryID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`withdrawal_history`), qm.WhereIn(`withdrawal_history.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load WithdrawalHistory")
	}

	var resultSlice []*WithdrawalHistory
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice WithdrawalHistory")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for withdrawal_history")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for withdrawal_history")
	}

	if len(withdrawalFiatAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.WithdrawalHistory = foreign
		if foreign.R == nil {
			foreign.R = &withdrawalHistoryR{}
		}
		foreign.R.WithdrawalFiats = append(foreign.R.WithdrawalFiats, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.WithdrawalHistoryID == foreign.ID {
				local.R.WithdrawalHistory = foreign
				if foreign.R == nil {
					foreign.R = &withdrawalHistoryR{}
				}
				foreign.R.WithdrawalFiats = append(foreign.R.WithdrawalFiats, local)
				break
			}
		}
	}

	return nil
}

// SetWithdrawalHistory of the withdrawalFiat to the related item.
// Sets o.R.WithdrawalHistory to related.
// Adds o to related.R.WithdrawalFiats.
func (o *WithdrawalFiat) SetWithdrawalHistory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *WithdrawalHistory) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"withdrawal_fiat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"withdrawal_history_id"}),
		strmangle.WhereClause("\"", "\"", 0, withdrawalFiatPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.WithdrawalHistoryID = related.ID
	if o.R == nil {
		o.R = &withdrawalFiatR{
			WithdrawalHistory: related,
		}
	} else {
		o.R.WithdrawalHistory = related
	}

	if related.R == nil {
		related.R = &withdrawalHistoryR{
			WithdrawalFiats: WithdrawalFiatSlice{o},
		}
	} else {
		related.R.WithdrawalFiats = append(related.R.WithdrawalFiats, o)
	}

	return nil
}

// WithdrawalFiats retrieves all the records using an executor.
func WithdrawalFiats(mods ...qm.QueryMod) withdrawalFiatQuery {
	mods = append(mods, qm.From("\"withdrawal_fiat\""))
	return withdrawalFiatQuery{NewQuery(mods...)}
}

// FindWithdrawalFiat retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWithdrawalFiat(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*WithdrawalFiat, error) {
	withdrawalFiatObj := &WithdrawalFiat{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"withdrawal_fiat\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, withdrawalFiatObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlite3: unable to select from withdrawal_fiat")
	}

	return withdrawalFiatObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WithdrawalFiat) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlite3: no withdrawal_fiat provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(withdrawalFiatColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	withdrawalFiatInsertCacheMut.RLock()
	cache, cached := withdrawalFiatInsertCache[key]
	withdrawalFiatInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			withdrawalFiatAllColumns,
			withdrawalFiatColumnsWithDefault,
			withdrawalFiatColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(withdrawalFiatType, withdrawalFiatMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(withdrawalFiatType, withdrawalFiatMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"withdrawal_fiat\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"withdrawal_fiat\" () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"withdrawal_fiat\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, withdrawalFiatPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to insert into withdrawal_fiat")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == withdrawalFiatMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to populate default values for withdrawal_fiat")
	}

CacheNoHooks:
	if !cached {
		withdrawalFiatInsertCacheMut.Lock()
		withdrawalFiatInsertCache[key] = cache
		withdrawalFiatInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the WithdrawalFiat.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WithdrawalFiat) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	withdrawalFiatUpdateCacheMut.RLock()
	cache, cached := withdrawalFiatUpdateCache[key]
	withdrawalFiatUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			withdrawalFiatAllColumns,
			withdrawalFiatPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("sqlite3: unable to update withdrawal_fiat, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"withdrawal_fiat\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, withdrawalFiatPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(withdrawalFiatType, withdrawalFiatMapping, append(wl, withdrawalFiatPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update withdrawal_fiat row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by update for withdrawal_fiat")
	}

	if !cached {
		withdrawalFiatUpdateCacheMut.Lock()
		withdrawalFiatUpdateCache[key] = cache
		withdrawalFiatUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q withdrawalFiatQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all for withdrawal_fiat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected for withdrawal_fiat")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WithdrawalFiatSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlite3: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), withdrawalFiatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"withdrawal_fiat\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, withdrawalFiatPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to update all in withdrawalFiat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to retrieve rows affected all in update all withdrawalFiat")
	}
	return rowsAff, nil
}

// Delete deletes a single WithdrawalFiat record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WithdrawalFiat) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlite3: no WithdrawalFiat provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), withdrawalFiatPrimaryKeyMapping)
	sql := "DELETE FROM \"withdrawal_fiat\" WHERE \"id\"=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete from withdrawal_fiat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by delete for withdrawal_fiat")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q withdrawalFiatQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlite3: no withdrawalFiatQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from withdrawal_fiat")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for withdrawal_fiat")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WithdrawalFiatSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(withdrawalFiatBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), withdrawalFiatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"withdrawal_fiat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, withdrawalFiatPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: unable to delete all from withdrawalFiat slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlite3: failed to get rows affected by deleteall for withdrawal_fiat")
	}

	if len(withdrawalFiatAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *WithdrawalFiat) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWithdrawalFiat(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WithdrawalFiatSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WithdrawalFiatSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), withdrawalFiatPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"withdrawal_fiat\".* FROM \"withdrawal_fiat\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, withdrawalFiatPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlite3: unable to reload all in WithdrawalFiatSlice")
	}

	*o = slice

	return nil
}

// WithdrawalFiatExists checks if the WithdrawalFiat row exists.
func WithdrawalFiatExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"withdrawal_fiat\" where \"id\"=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlite3: unable to check if withdrawal_fiat exists")
	}

	return exists, nil
}
