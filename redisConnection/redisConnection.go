package redisconnection

import (
	"github.com/garyburd/redigo/redis"
)

func OpenConnection() redis.Conn {

	var err error

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	return c
}

func CloseConnection(c redis.Conn) {
	c.Close()
}

func SaveMessage(c redis.Conn) {
	c.Do("SET", "message1", "Hello World")
}
