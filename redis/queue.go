package redis

import (
	"github.com/hoisie/redis"
)

// get instance of Queue
func NewQueue(name string) *Queue {
	redisClient := GetRedisClientFromConfig()
	queue := &Queue{Name: name, RedisClient: redisClient}
	return queue
}

type Queue struct {
	Name        string
	RedisClient *redis.Client
}

//push queue
func (this *Queue) Push(input string) bool {
	err := this.RedisClient.Lpush(this.Name, []byte(input))
	if err != nil {
		return false
	} else {
		return true
	}
}

//pop queue
func (this *Queue) Pop() (string, error) {
	sd, err := this.RedisClient.Rpop(this.Name)
	if err == nil {
		return string(sd), nil
	} else {
		return "", err
	}
}
