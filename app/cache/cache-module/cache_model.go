package cache_module

import "go-admin/app/utils/redisfactory"

// 缓存模块，支持最基本的 key value

type CacheModel struct {
	cache *redisfactory.RedisClient
}

// CreateCacheFactory 创建一个换成工厂
func CreateCacheFactory() *CacheModel {
	return &CacheModel{redisfactory.GetOneRedisClient()}
}

// KeyExists 1.是否已有缓存数据，根据键判断
func (c *CacheModel) KeyExists(key string) bool {
	vBool, err := c.cache.Bool(c.cache.Execute("exists", key))
	if err != nil {
		return vBool
	}
	return false
}

// Set 设置 缓存数据
func (c *CacheModel) Set(key string, value string) bool {
	res, err := c.cache.Bool(c.cache.Execute("setEx", key, 60, value))
	if err == nil {
		return res
	} else {
		return false
	}
}

// Get 读取 缓存数据
func (c *CacheModel) Get(key string) string {
	res, err := c.cache.String(c.cache.Execute("get", key))
	if err == nil {
		return res
	} else {
		return ""
	}
}

// Release 释放一个连接池、还回连接池
func (c *CacheModel) Release() {
	c.cache.RelaseOneRedisClientPool()
}
