package dao

import (
	"errors"
	"fmt"
	cfgloader "myweb/config"
	"myweb/db"
)

// 每个XXXDao包通过init函数注册函数值(写入daoRegister中)
var daoRegister map[string]func(c *db.Connection)

// LoadConnections 装载数据库连接.
// 遍历配置文件中的db表, 并创建相应的数据库连接. 用db表中的key在daoRegister中查找, 如果存在, 则将创建的数据库连接传入函数值进行处理。
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
