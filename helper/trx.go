package helper

import (
	"fmt"
	"gorm.io/gorm"
)

func CommitOrRollback(trx *gorm.DB) {
	err := recover()
	if err != nil {
		trx.Rollback()
		fmt.Println("transaction rollback")
	} else {
		trx.Commit()
		fmt.Println("transaction commit")
	}
}
