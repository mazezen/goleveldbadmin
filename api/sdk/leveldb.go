package sdk

import (
	errors2 "errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	ErrEmptyKey = errors.New("key could not be empty")
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
	}
}
