package helper

import (
	"github.com/spf13/viper"
	"go-admin/app/global/variable"
	"log"
	"time"
)

func CreateYamlFactory() *ConfigYml {
	yamlConfig := viper.New()
	//windows环境下为 %GOPATH，linux环境下为 $GOPATH
	yamlConfig.AddConfigPath(variable.BASE_PATH + "/config")
	// 需要读取的文件名
	yamlConfig.SetConfigName("config")
	// 设置配置文件类型
	yamlConfig.SetConfigType("yaml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatalf("初始化配置文件发生错误：%s\n", err)
	}

	return &ConfigYml{
		yamlConfig,
	}
}

// ConfigYml yaml配置类
type ConfigYml struct {
	viper *viper.Viper
}

// Get get一个原始值
func (c *ConfigYml) Get(keyName string) interface{} {
	return c.viper.Get(keyName)
}

// GetString getString
func (c *ConfigYml) GetString(keyName string) string {
	return c.viper.GetString(keyName)
}

// GetBool get bool
func (c *ConfigYml) GetBool(keyName string) bool {
	return c.viper.GetBool(keyName)
}

// GetInt getting
func (c *ConfigYml) GetInt(keyName string) int {
	return c.viper.GetInt(keyName)
}

// GetInt32 getting32
func (c *ConfigYml) GetInt32(keyname string) int32 {
	return c.viper.GetInt32(keyname)
}

// GetInt64 getting64
func (c *ConfigYml) GetInt64(keyname string) int64 {
	return c.viper.GetInt64(keyname)
}

// GetDuration getDuration
func (c *ConfigYml) GetDuration(keyname string) time.Duration {
	return c.viper.GetDuration(keyname)
}
