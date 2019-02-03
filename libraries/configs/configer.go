package configs

import (
	"errors"
	"reflect"
	"time"

	config "github.com/zpatrick/go-config"
)

// Configer represeting the config interface of .ini file
type Configer interface {
	GetPrefix() string
	SetPrefix(prefix string)
	SetPath(path string) error
	GetPath() string
}

// Config representing basic config parameters
type Config struct {
	Prefix string `json:"-"`
	Path   string `json:"-"`
}

// GetPrefix returns the current config prefix value
func (cfg *Config) GetPrefix() string {
	return cfg.Prefix
}

// SetPrefix sets the config prefix value
func (cfg *Config) SetPrefix(prefix string) {
	cfg.Prefix = prefix
}

// SetPath sets the path of the .ini file
func (cfg *Config) SetPath(path string) error {
	cfg.Path = path
	return nil
}

// GetPath returns the path of the .ini file
func (cfg *Config) GetPath() string {
	return cfg.Path
}

// LoadConfig loads the .ini file in specific path
func LoadConfig(cfg Configer) error {
	iniFile := config.NewINIFile(cfg.GetPath())
	c := config.NewConfig([]config.Provider{iniFile})
	if err := c.Load(); err != nil {
		return err
	}

	return parseConfig(cfg, c)
}

func parseConfig(cfg Configer, c *config.Config) error {
	if cfg.GetPrefix() == "" {
		return errors.New("Prefix is required")
	}
	configStruct := reflect.ValueOf(cfg).Elem()
	if !configStruct.IsValid() {
		return errors.New("fail to parse configuration data")
	}

	for i := 0; i < configStruct.NumField(); i++ {
		field := configStruct.Field(i)
		name := configStruct.Type().Field(i).Tag.Get("json")
		if name != "-" {
			// property format is like local.port
			property := cfg.GetPrefix() + "." + name
			switch field.Kind() {
			case reflect.Int:
				v, err := c.Int(property)
				if err != nil {
					return err
				}
				field.SetInt(int64(v))
			case reflect.Float64:
				v, err := c.Float(property)
				if err != nil {
					return err
				}
				field.SetFloat(v)
			case reflect.String:
				v, err := c.String(property)
				if err != nil {
					return err
				}
				field.SetString(v)
			case reflect.Bool:
				v, err := c.Bool(property)
				if err != nil {
					return err
				}
				field.SetBool(v)
			case reflect.Ptr:
				v, err := c.String(property)
				if err != nil {
					return err
				}
				ptrValue, err := time.LoadLocation(v)
				if err != nil {
					return err
				}
				rfValue := reflect.ValueOf(ptrValue)
				field.Set(rfValue)
			}
		}
	}
	return nil
}
