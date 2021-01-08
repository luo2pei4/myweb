package dao

import (
	"errors"
	"fmt"
	cfgloader "myweb/config"
	"myweb/db"
)

var daoRegister map[string]func(c *db.Connection)

// LoadConnections 装载数据库连接
func LoadConnections() error {

	cfgs, err := cfgloader.GetTable("db")

	if err != nil {
		return err
	}

	keys := cfgs.Keys()

	if keys == nil || len(keys) == 0 {
		return errors.New("No configuration")
	}

	for _, key := range keys {

		driver := cfgs.Get(key + ".driver")
		dsn := cfgs.Get(key + ".dsn")

		conn, err := db.NewConnection(key, driver.(string), dsn.(string))

		if err != nil {
			msg := fmt.Sprintf("Create %s connection failed. Details:\n %s", key, err.Error())
			return errors.New(msg)
		}

		f := daoRegister[key]
		f(conn)
	}

	return nil
}

// RegistFunc 注册方法
func RegistFunc(key string, f func(conn *db.Connection)) {

	if daoRegister == nil {
		daoRegister = make(map[string]func(conn *db.Connection))
	}

	daoRegister[key] = f
}
