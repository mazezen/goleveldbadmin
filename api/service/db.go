package service

import (
	"encoding/json"
	"errors"
	"github.com/mazezen/goleveldbadmin/sdk"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type DbService struct{}

type LRes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// FindAll 数据中心
func (d *DbService) FindAll(page, pageSize int, key string) (int64, []LRes) {
	var r = make([]LRes, 0)

	var total int64
	i := sdk.Db.NewIterator(nil, nil)
	defer i.Release()
	for i.Next() {
		total++
	}

	if len(key) > 0 {
		i2 := sdk.Db.NewIterator(nil, nil)
		defer i2.Release()
		for i2.Next() {
			key1 := string(i2.Key())
			value := i2.Value()
			if key1 == key {
				res := LRes{
					Key:   key,
					Value: string(value),
				}
				r = append(r, res)
			}
		}
		return total, r
	}

	var startKey []byte
	iterator := sdk.Db.NewIterator(nil, nil)
	defer iterator.Release()
	var ps int64
	for iterator.Next() {
		if page == 1 {
			startKey = iterator.Key()
			break
		}
		ps++
		if ps == int64((page-1)*pageSize+1) {
			startKey = iterator.Key()
			break
		}
	}

	newIterator := sdk.Db.NewIterator(&util.Range{Start: startKey}, nil)
	defer newIterator.Release()
	count := 0
	for newIterator.Next() {
		key := string(newIterator.Key())
		value := string(newIterator.Value())
		if count >= pageSize {
			break
		}
		res := LRes{
			Key:   key,
			Value: value,
		}
		count++
		r = append(r, res)
	}

	return total, r
}

// Add 添加
func (d *DbService) Add(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("empty key")
	}
	res, _ := json.Marshal(value)
	err := sdk.Db.Put([]byte(key), res, nil)
	if err != nil {
		return err
	}
	return nil
}

// Info 详情
func (d *DbService) Info(key string) (*LRes, error) {
	if len(key) == 0 {
		return nil, errors.New("empty key")
	}
	value, err := sdk.Db.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}
	res := &LRes{
		Key:   key,
		Value: string(value),
	}
	return res, nil
}

// Delete 删除
func (d *DbService) Delete(key string) error {
	if len(key) == 0 {
		return errors.New("empty key")
	}
	return sdk.Db.Delete([]byte(key), nil)
}

// Clear 清空
func (d *DbService) Clear() error {
	iterator := sdk.Db.NewIterator(nil, nil)
	defer iterator.Release()
	for iterator.Next() {
		key := iterator.Key()
		if err := d.Delete(string(key)); err != nil {
			return err
		}
	}
	return nil
}
