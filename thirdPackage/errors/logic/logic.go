package logic

import (
	"myGoPractices/thirdPackage/errors/dao"

	"github.com/pkg/errors"
)

func ReturnErr() error {
	err := dao.ReturnErr()
	return errors.Wrap(err, "logic error")
}
