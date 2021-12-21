package configservice

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"strconv"
	"sync"

	"github.com/spf13/viper"
	"github.com/tim5wang/selfman/common/util"
)

type KVEngine interface {
	GetString(key string) string
}

type FileConfig interface {
	SetFS(fs fs.FS) error
}

type YamlConfig struct {
	file     string
	path     string
	fileType string
	viper    *viper.Viper
}

func (y *YamlConfig) fullName() string {
	return fmt.Sprintf("%s.%s", path.Join(y.path, y.file), y.fileType)
}

func NewYamlConfig(p, file, fileType string) KVEngine {
	c := &YamlConfig{
		path:     p,
		file:     file,
		fileType: fileType,
	}
	c.viper = viper.New()
	f := c.fullName()
	c.viper.SetConfigType(fileType)
	c.viper.SetConfigName(file)
	c.viper.SetConfigFile(f)
	_ = c.viper.ReadInConfig()
	//allConfig := c.viper.AllSettings()
	//util.Print(allConfig)
	return c
}

func (c *YamlConfig) SetFS(fs fs.FS) error {

	file, err := fs.Open(c.fullName())
	if err != nil {
		return err
	}
	c.viper.SetConfigName(c.file)
	c.viper.SetConfigType(c.fileType)
	err = c.viper.ReadConfig(file)
	if err != nil {
		return err
	}
	//allConfig := c.viper.AllSettings()
	//util.Print("allconfig:", allConfig)
	return nil
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

func (c *ConfigService) FromEmbed(e fs.FS) error {
	for _, kv := range c.opts.Engines {
		fileConfig, ok := kv.(FileConfig)
		if ok {
			err := fileConfig.SetFS(e)
			if err != nil {
				util.Print(err)
			}
		}
	}
	return nil
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
