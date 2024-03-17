package config

import (
	"database/sql"
	"sync"
)

type ServerDb struct {
	Db *sql.DB
}

var instance *ServerDb
var once sync.Once

func GetInstance() *ServerDb {
	once.Do(func() {
		db, err := sql.Open("sqlite3", "./"+GetEnv(ServerDbName))
		if err != nil {
			panic(err.Error())
		}
		instance = &ServerDb{Db: db}
	})
	return instance
}
