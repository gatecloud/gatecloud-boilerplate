package configs

import (
	"errors"
	"log"
	"path"
	"runtime"

	env "gatecloud-boilerplate/libraries/configs"
)

// LocalConfig stores the configuration of the local environment variable
type LocalConfig struct {
	env.Config
	Port          string   `json:"port"`
	Production    bool     `json:"production"`
	AutoMigration bool     `json:"auto_migration"`
	CorsEnabled   bool     `json:"cors_enabled"`
	Env           *env.Env `json:"-"`
}

// GetPrefix returns the prefix in .ini file
func (lcfg *LocalConfig) GetPrefix() string {
	return lcfg.Prefix
}

// SetPrefix sets the prefix in .ini file
func (lcfg *LocalConfig) SetPrefix(prefix string) error {
	if lcfg == nil {
		return errors.New("The Config object is not initialized")
	}
	lcfg.Prefix = prefix
	return nil
}

// SetPath sets the path of the target file
func (lcfg *LocalConfig) SetPath(path string) error {
	if lcfg == nil {
		return errors.New("The Config object is not initialized")
	}
	if path == "" {
	}

	lcfg.Path = path
	return nil
}

// GetPath returns the path of the target file
func (lcfg *LocalConfig) GetPath() string {
	return lcfg.Path
}

// Configuration is proxy's global configuration
var Configuration *LocalConfig

func init() {
	Configuration = &LocalConfig{
		Env: &env.Env{},
	}

	// Load the env.ini file
	if err := Configuration.Env.SetPath(""); err != nil {
		log.Fatal(err)
	}

	if err := Configuration.Env.SetPrefix("env"); err != nil {
		log.Fatal(err)
	}

	if err := env.LoadConfig(Configuration.Env); err != nil {
		log.Fatal(err)
	}

	// Load the local.ini file
	_, dir, _, _ := runtime.Caller(1)
	if err := Configuration.SetPath(path.Join(path.Dir(dir), "/configs/local.ini")); err != nil {
		log.Fatal(err)
	}

	if err := Configuration.SetPrefix("api"); err != nil {
		log.Fatal(err)
	}

	if err := env.LoadConfig(Configuration); err != nil {
		log.Fatal(err)
	}

}
