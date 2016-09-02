package daos

import (
	"github.com/garyburd/redigo/redis"
)

type RedisDao struct {
}

func (c *RedisDao) GetIdByReids(name string) (int64, error) {
	pool := Pool.Get()
	defer pool.Close()

	id, err := redis.Int64(pool.Do("GET", TableName+":"+name))
	if err != nil || id <= 0 {
		pool.Do("SET", TableName+":"+name, StartNum)
	}

	newId, err := redis.Int64(pool.Do("INCR", TableName+":"+name))

	return newId, err
}
