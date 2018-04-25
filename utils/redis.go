package utils

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"shuaibingBlog/helper"
	"time"
)

var Conn redis.Conn
var Cache Redis

type Redis struct {
	conn redis.Conn
}

func init() {

	//使用redis封装的Dial进行tcp连接
	c, err := redis.Dial("tcp", helper.Config.Redisconn)
	if err != nil {
		panic("redis err" + helper.Config.Redisconn)
	}
	Conn = c
	Cache = Redis{conn: c}
}

//Get 获取一个值
func (r *Redis) Get(key string) interface{} {
	conn := r.conn

	var data []byte
	var err error

	if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
		return nil
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}

	return reply
}

//Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r.conn

	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}

	_, err = conn.Do("SETEX", key, int64(timeout/time.Second), data)

	return
}

//IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	conn := r.conn
	a, _ := conn.Do("EXISTS", key)
	i := a.(int64)
	if i > 0 {
		return true
	}
	return false
}

//Delete 删除
func (r *Redis) Delete(key string) error {
	conn := r.conn

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}
