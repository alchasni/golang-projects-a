package service

import "strings"

type Error interface {
	error
	Code() string
}

type serviceError struct {
	code    string
	message string
}

var _ Error = serviceError{}

func (s serviceError) Error() string {
	return s.message
}

func (s serviceError) Code() string {
	return s.code
}

const (
	ErrorCode_General          = "system:common:general"
	ErrorCode_InvalidInput     = "system:common:invalid-input"
	ErrorCode_DatasourceAccess = "system:common:datasource-access"
)

var (
	NewErr = func(code, msg string, msgs ...string) Error {
		if len(msgs) > 0 {
			msg = strings.Join(msgs, "; ")
		}
		return serviceError{
			code:    code,
			message: msg,
		}
	}

	ErrGeneral = func(msgs ...string) Error {
		return NewErr(ErrorCode_General, "general error", msgs...)
	}

	ErrInvalidInput = func(msgs ...string) Error {
		return NewErr(ErrorCode_InvalidInput, "invalid input", msgs...)
	}

	ErrDatasourceAccess = func(msgs ...string) Error {
		return NewErr(ErrorCode_DatasourceAccess, "datasource access error", msgs...)
	}
)
