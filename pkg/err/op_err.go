package slerr

import (
	"fmt"
	"github.com/defany/slogger/pkg/logger/sl"
)

func WithSource(err error) error {
	op := sl.FnName(1)

	return fmt.Errorf("%s: %w", op, err)
}
