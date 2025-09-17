package slerr

import (
	"fmt"
	"strings"

	"github.com/defany/slogger/pkg/logger/sl"
)

func WithSource(err error, comments ...string) error {
	if err == nil {
		return nil
	}

	op := sl.FnName(1)

	if len(comments) > 0 {
		return fmt.Errorf("%s [%s] -> %w", op, strings.Join(comments, " "), err)
	}

	return fmt.Errorf("%s -> %w", op, err)
}
