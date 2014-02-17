package redis

import (
	"github.com/hoisie/redis"
)

// get instance of Set
func NewSet(name string) *Set {
	redisClient := GetRedisClientFromConfig()
	redisClient.Expire(name, 604800)
	set := &Set{Name: name, RedisClient: redisClient}
	return set
}

type Set struct {
	Name        string
	RedisClient *redis.Client
}

//add cache
func (this *Set) Add(input string) bool {
	ok, _ := this.RedisClient.Sadd(this.Name, []byte(input))
	if ok {
		return true
	} else {
		return false
	}
}

//get cache
func (this *Set) IsMember(input string) bool {
	ok, _ := this.RedisClient.Sismember(this.Name, []byte(input))
	if ok {
		return true
	} else {
		return false
	}
}
