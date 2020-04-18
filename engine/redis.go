package engine

import "github.com/gomodule/redigo/redis"

// Dial : open connection to redis
func Dial() (redis.Conn, error) {
	c, err := redis.Dial("tcp", ":6379")
	return c, err
}

// AddFile : add an entry for a hit in redis
func AddFile(c redis.Conn, key, file string, score int) error {
	_, err := c.Do("ZADD", key, score, file)
	return err
}

// Get : lookup an entry for a hit in redis
func Get(c redis.Conn, key string) ([]string, error) {
	res, err := redis.Strings(c.Do("ZREVRANGE", key, 0, -1))
	return res, err
}
