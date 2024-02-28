package redisfactory

import (
	"github.com/gomodule/redigo/redis"
	"go-admin/app/global/errors"
	"go-admin/app/utils/helper"
	"log"
	"time"
)

// 创建redis连接池
func createRedisClientPool() *redis.Pool {
	configFac := helper.CreateYamlFactory()

	return &redis.Pool{
		MaxIdle:     configFac.GetInt("redis.maxIdle"),                        //最大空闲数
		MaxActive:   configFac.GetInt("redis.maxActive"),                      //最大活跃数
		IdleTimeout: configFac.GetDuration("redis.idleTimeout") * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", configFac.GetString("redis.host")+":"+configFac.GetString("redis.port"))
			if err != nil {
				log.Fatal(errors.Errors_Redis_InitConnFail, err)
			}
			conn.Do("select", configFac.GetInt("redis.indexDb"))
			password := configFac.GetString("redis.password") // 通过配置选项设置redis密码
			if len(password) >= 5 {
				if _, err := conn.Do("AUTH", password); err != nil {
					defer conn.Close()
					log.Fatal(errors.Errors_Redis_AuhtFail, err)
				}
			}
			return conn, err
		},
	}
}

// GetOneRedisClient 从连接池中获取一个redis连接
func GetOneRedisClient() *RedisClient {
	poolConn := createRedisClientPool()
	return &RedisClient{client: poolConn.Get()}
}

// RedisClient 定义一个redis客户端结构体
type RedisClient struct {
	client redis.Conn
}

// Execute 为redis-go 客户端封装统一操作函数入口
func (r *RedisClient) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return r.client.Do(cmd, args)
}

// RelaseOneRedisClientPool 释放连接池
func (r *RedisClient) RelaseOneRedisClientPool() {
	r.client.Close()
}

//  封装几个数据类型转换的函数

// Bool bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

//string 类型转换
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// Strings strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

//Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

// Int int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

// Int64 int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

// Uint64 uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

//Bytes 类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}
