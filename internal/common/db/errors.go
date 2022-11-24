package db

import "errors"

var NoDataWasInsertedOrUpdatedError = errors.New("no data was inserted or updated")

var NoRowsInResultErrorMessage = "sql: no rows in result set"
