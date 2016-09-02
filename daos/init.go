package daos

import (
	"fmt"
	"idgo/logger"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool
var StartNum int64
var TableName string

func init() {
	redisInit()
}

func redisInit() {
	StartNum, _ = beego.AppConfig.Int64("startNum")
	TableName = beego.AppConfig.String("appname")
	mode := beego.BConfig.RunMode
	host := beego.AppConfig.String(mode + "::redisHost")
	passwd := beego.AppConfig.String(mode + "::redisPasswd")
	pool := &redis.Pool{
		MaxIdle: 100,
		//	MaxActive:   2000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				logger.ErrLogs.Error(err.Error())
				return nil, err
			}
			if _, err = c.Do("AUTH", passwd); err != nil {
				logger.ErrLogs.Error(err.Error())
				c.Close()
				return nil, err
			}

			_, err = c.Do("SELECT", "1")
			if err != nil {
				logger.ErrLogs.Error(err.Error())
				return nil, err
			}

			return c, err
		},
	}
	conn := pool.Get()
	if conn.Err() != nil {
		logger.ErrLogs.Error(conn.Err().Error())
		fmt.Println("redis数据库连接失败")
		os.Exit(-1)
	}
	Pool = pool
}
