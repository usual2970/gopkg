package sql

import (
	sqErr "github.com/usual2970/gopkg/errors"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func FinishTransaction(err error, tx *gorm.DB) error {

	if err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return sqErr.NewAggregate([]error{err, rbErr})
		}
		return err
	} else {
		if err := tx.Commit().Error; err != nil {
			return errors.Wrap(err, "failed to commit tx")
		}
		return nil
	}
}
