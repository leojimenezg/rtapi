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
	return fmt.Sprintf("failed to iterate on all query rows:\n%v", e.Err)
}

type FillColumnsError struct {
	Err error
}
func (e FillColumnsError) Error() string {
	return fmt.Sprintf("failed to fill Category struct with row columns:\n%v", e.Err)
}
