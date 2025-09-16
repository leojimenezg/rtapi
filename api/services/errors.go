package services

import (
	"fmt"
)

type BadQueryError struct {
	Query string
	Err error
}
func (e BadQueryError) Error() string {
	return fmt.Sprintf("failed to query to the database %s:\n%v", e.Query, e.Err)
}

type RowsIterationError struct {
	Query string
	Err error
}
func (e RowsIterationError) Error() string {
	return fmt.Sprintf("failed to iterate on all rows from %s query:\n%v", e.Query, e.Err)
}

type FillColumnsError struct {
	Err error
}
func (e FillColumnsError) Error() string {
	return fmt.Sprintf("failed to fill Category struct with row columns:\n%v", e.Err)
}

type UpdateRegisterError struct {
	Query string
	Err error
}
func (e UpdateRegisterError) Error() string {
	return fmt.Sprintf("failed to update register using %s query:\n%v", e.Query, e.Err)
}

type NotFoundError struct {
	Query string
	Field string
	Value any
}
func (e NotFoundError) Error() string {
	return fmt.Sprintf("no rows found using %s query and %s = %v", e.Query, e.Field, e.Value)
}
