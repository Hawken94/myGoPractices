package dao

import (
	"github.com/pkg/errors"
)

func ReturnErr() error {
	return errors.WithStack(errors.New("message string"))
}
