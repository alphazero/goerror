package goerror

import (
	"errors"
)

// ideally we want optional args capped at 1 item
// but you can't do that in go.
type errFn func(...string) error

// need to export this to surface Is() to docs
type ErrPredicate struct{ error }

const errPrefix = "error - "
const errPrefixLen = len(errPrefix)

// defines a new categorical error.
func Define(category string) errFn {
	return func(args ...string) error {
		errstr := errPrefix + category
		if len(args) == 0 {
			return errors.New(errstr)
		}
		errstr += " - "
		// since nargs can be > 1 might as well
		// make a virtue of it and pretty concat the args
		for _, arg := range args {
			errstr += arg
			errstr += " "
		}
		errstr = errstr[:len(errstr)-1]
		return errors.New(errstr)
	}
}

// Returns an ErrPredicate, typically for use in conjunction
// with the ErrPredicate#Is(). Function name is as such to
// allow for a readable call site, as below:
//
//     if errors.TypeOf(e).Is(AssertionError)
//
func TypeOf(e error) *ErrPredicate {
	return &ErrPredicate{e}
}

// Returns true if the ErrPredicate.error is an 'instance'
// of input arg 'efn'.
func (e *ErrPredicate) Is(efn errFn) bool {
	s := e.Error()
	category := efn().Error()
	catlen := len(category)
	if len(s) < catlen || s[:catlen] != category {
		return false
	}

	return true
}
