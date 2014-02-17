package redis

import (
	"github.com/hoisie/redis"
	. "github.com/tabalt/appgo"
)

//get redis client
func GetRedisClient(host string, port string, password string, database int) * redis.Client {
	addr := host + ":" + port
    redisClient := &redis.Client{Addr:addr, Db: database, Password: password}
    //redisClient.Expire(cacheName, expire)
    return redisClient
}

//get redis client from config
func GetRedisClientFromConfig() * redis.Client {
	
	host, _ := AppConfig.String("redis_host")
	port, _ := AppConfig.String("redis_port")
    password, _ := AppConfig.String("redis_password")
    //database, _ := AppConfig.Int("redis_database")
    //TODO get int from config
    database := 0
    return GetRedisClient(host, port, password, database)
}