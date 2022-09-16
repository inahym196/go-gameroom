package model_redis

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func NewSession(c *gin.Context, cookieKey, redisValue string) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Fatal(err)
	}
	newRedisKey := base64.URLEncoding.EncodeToString(b)
	if err := conn.Set(c, newRedisKey, redisValue, 0).Err(); err != nil {
		log.Fatal(err)
	}
	c.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
}

func GetSession(c *gin.Context, cookieKey string) interface{} {
	redisKey, _ := c.Cookie(cookieKey)
	redisValue, err := conn.Get(c, redisKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("SessionKeyが登録されていません")
		return nil
	case err != nil:
		fmt.Println("Session取得時にエラー発生: " + err.Error())
		return nil
	}
	return redisValue
}

func DeleteSession(c *gin.Context, cookieKey string) {
	redisId, _ := c.Cookie(cookieKey)
	conn.Del(c, redisId)
	c.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)
}
