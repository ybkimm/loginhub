package app

import "errors"

type WrappedError struct {
	msg string
	err error
}

func wrapErr(err error, msg string) error {
	return &WrappedError{
		msg: msg,
		err: err,
	}
}

func (e *WrappedError) Unwrap() error {
	return e.err
}

func (e *WrappedError) Error() string {
	return "app: " + e.msg + ": " + e.err.Error()
}

var ErrConfigNotLoaded = errors.New("app: configuration not loaded")
