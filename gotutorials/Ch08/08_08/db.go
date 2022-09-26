// Code from https://github.com/LinkedInLearning/go-essential-training-2446262/blob/main/Ch08/08_08/db.go
package main

import "sync"

type DB struct {
	m sync.Map
}

func (db *DB) Get(key string) []byte {
	val, ok := db.m.Load(key)
	if !ok {
		return nil
	}

	return val.([]byte)
}

func (db *DB) Set(key string, value []byte) {
	db.m.Store(key, value)
}
