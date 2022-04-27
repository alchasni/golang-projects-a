package mysql

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

const (
	errorCode_ER_DUP_ENTRY uint16 = 1062
)

func errorIs(err error, errCode uint16) bool {
	var mysqlErr interface{}
	return errors.As(err, &mysqlErr) && mysqlErr.(mysql.MySQLError).Number == errCode
}
