package tokens

import "errors"

var (
	ErrInvalidToken = errors.New("tokens: invalid token")
	ErrExpired      = errors.New("tokens: expired")
	ErrRevoked      = errors.New("tokens: revoked")
)

type ErrWrapper struct {
	err error
}

func wrapErr(e error) *ErrWrapper {
	return &ErrWrapper{e}
}

func (e *ErrWrapper) Unwrap() error {
	return e.err
}

func (e *ErrWrapper) Error() string {
	return "tokens: " + e.err.Error()
}
