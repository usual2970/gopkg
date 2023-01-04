package gorm

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/usual2970/gopkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db atomic.Value

func GetDB() (*gorm.DB, error) {
	ld := db.Load()
	if ld == nil {
		dbHost := conf.GetString(`database.host`)
		dbPort := conf.GetString(`database.port`)
		dbUser := conf.GetString(`database.user`)
		dbPass := conf.GetString(`database.pass`)
		dbName := conf.GetString(`database.name`)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
		c, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		sqlDb, err := c.DB()
		if err != nil {
			return nil, err
		}
		sqlDb.SetMaxIdleConns(conf.GetInt(`database.max_idle_conns`))
		sqlDb.SetMaxOpenConns(conf.GetInt(`database.max_open_conns`))

		db.Store(c)
		return c, nil
	}

	rs, ok := ld.(*gorm.DB)

	if ok {
		return rs, nil
	}

	return nil, errors.New("db is not a gorm.DB")
}
