package configservice

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/spf13/viper"
	"github.com/tim5wang/selfman/common/util"
)

type KVEngine interface {
	GetString(key string) string
}

type YamlConfig struct {
	file  string
	viper *viper.Viper
}

func NewYamlConfig(file string) KVEngine {
	c := &YamlConfig{
		file: file,
	}
	c.viper = viper.New()
	c.viper.SetConfigFile(file)
	c.viper.SetConfigType("yaml")
	err := c.viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	allConfig := c.viper.AllSettings()
	util.PrintJSONIndent(allConfig)
	return c
}

func (c *YamlConfig) GetString(key string) (value string) {
	value = c.viper.GetString(key)
	return
}

type config interface {
	KVEngine
	GetBool(key string) bool
	GetInt(key string) int
	GetInt64(key string) int64
	GetStruct(key string, value interface{})
}

type Options struct {
	Engines   []KVEngine
	Path      string
	WithCache bool
}

type ConfigService struct {
	opts   *Options
	lock   sync.RWMutex
	config map[string]interface{}
}

func NewConfigService(op *Options) *ConfigService {
	c := &ConfigService{
		opts:   op,
		config: make(map[string]interface{}),
	}
	return c
}

func (c *ConfigService) GetString(key string) (value string) {
	if c.opts.WithCache {
		lock := c.lock.RLocker()
		lock.Lock()
		if v, ok := c.config[key]; ok {
			value, _ = v.(string)
			return
		}
		lock.Unlock()
	}
	for _, reader := range c.opts.Engines {
		value = reader.GetString(key)
		if value != "" {
			c.lock.Lock()
			c.config[key] = value
			c.lock.Unlock()
			return
		}
	}
	return
}
func (c *ConfigService) GetBool(key string) (value bool) {
	v := c.GetString(key)
	if v != "" {
		value = v == "true" || v == "True" || v == "TRUE" || v == "1"
	}
	return
}
func (c *ConfigService) GetInt(key string) (value int) {
	v := c.GetInt64(key)
	value = int(v)
	return
}
func (c *ConfigService) GetInt64(key string) (value int64) {
	v := c.GetString(key)
	if v != "" {
		vint64, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			value = vint64
		}
	}
	return
}
func (c *ConfigService) GetStruct(key string, value interface{}) {
	v := c.GetString(key)
	if v != "" {
		_ = json.Unmarshal([]byte(v), value)
	}
}

func (c *ConfigService) Hock(key, value string) (ok bool) {
	c.config[key] = value
	ok = true
	return
}
