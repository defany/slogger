package slerr

import (
	"fmt"
	"github.com/defany/slogger/pkg/logger/sl"
	"strings"
)

func WithSource(err error, comments ...string) error {
	op := sl.FnName(1)

	if len(comments) > 0 {
		return fmt.Errorf("%s [%s] -> %w", op, strings.Join(comments, " "), err)
	}

	return fmt.Errorf("%s -> %w", op, err)
}
