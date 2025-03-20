package sdk

import (
	errors2 "errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"log"
)

var (
	err error
	Db  *leveldb.DB
)

func CreateLevelDB(path string) {
	Db, err = leveldb.OpenFile(path, nil)
	var errCorrupted *errors.ErrCorrupted
	if errors2.As(err, &errCorrupted) {
		Db, err = leveldb.RecoverFile(path, nil)
		if err != nil {
			log.Fatalf("leveldb recovery failed: %v", err)
		}
	}
}
