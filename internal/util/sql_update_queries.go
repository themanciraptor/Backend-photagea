package util

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// UpdateQueryBuilder allows us to build a safe dynamically created update string
type UpdateQueryBuilder struct {
	table   string
	args    []sql.NamedArg
	filters []sql.NamedArg
	errors  []error
}

// ErrNoFilters is returns when we attempt to execute a query with no filters
var ErrNoFilters = errors.New("UpdateQueryBuilder: no filters assigned to query")

// ErrNoUpdateFields is returns when we attempt to execute a query with no filters
var ErrNoUpdateFields = errors.New("UpdateQueryBuilder: no fields set to be updated by query")

// InitUpdateQueryBuilder initializes a query builder with information necessary to complete an update query
func InitUpdateQueryBuilder(tableName string) *UpdateQueryBuilder {
	return &UpdateQueryBuilder{table: tableName}
}

// Add another named update argument
func (u *UpdateQueryBuilder) Add(field string, value interface{}) *UpdateQueryBuilder {

	if reflect.ValueOf(value).IsZero() {
		u.errors = append(u.errors, fmt.Errorf("QueryBuilder: statement value is null"))
	}
	if field == "" {
		u.errors = append(u.errors, fmt.Errorf("QueryBuilder: statement field is null"))
	}

	if len(u.errors) > 0 {
		return u
	}

	u.args = append(u.args, sql.Named(field, value))

	return u
}

// AddFilter adds a filter to deside what row to update
func (u *UpdateQueryBuilder) AddFilter(filterName string, value interface{}) *UpdateQueryBuilder {
	if reflect.ValueOf(value).IsZero() {
		u.errors = append(u.errors, fmt.Errorf("QueryBuilder: statement value is null"))
	}
	if filterName == "" {
		u.errors = append(u.errors, fmt.Errorf("QueryBuilder: statement field is null"))
	}

	if len(u.errors) > 0 {
		return u
	}

	u.filters = append(u.filters, sql.Named(filterName, value))

	return u
}

// BuildStatements assembles and returns an update string
func BuildStatements(args []sql.NamedArg) string {
	var statements []string
	for _, arg := range args {
		statements = append(statements, fmt.Sprintf("%s=?", arg.Name))
	}
	return strings.Join(statements, ",")
}

// GetErrors returns a list of all the errors that occurred while adding update statements
func (u *UpdateQueryBuilder) GetErrors() []error {
	return u.errors
}

// BuildUpdateQuery builds an update query
func (u *UpdateQueryBuilder) BuildUpdateQuery() (string, error) {
	if len(u.args) == 0 {
		return "", ErrNoUpdateFields
	}

	if len(u.filters) == 0 {
		return "", ErrNoFilters
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s", u.table, BuildStatements(u.args), BuildStatements(u.filters)), nil
}

// ExecuteQuery builds and executes an update query
func (u *UpdateQueryBuilder) ExecuteQuery(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q, err := u.BuildUpdateQuery()
	if err != nil {
		return nil, err
	}

	s := append(u.args, u.filters...)
	var args []interface{} = make([]interface{}, len(s))
	for i, arg := range s {
		args[i] = arg.Value
	}

	return db.ExecContext(ctx, q, args...)
}
