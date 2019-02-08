package configs

import (
	"log"
	"path"
	"runtime"

	libConfig "github.com/gatecloud/webservice-library/config"
)

// LocalConfig stores the configuration of the local environment variable
type LocalConfig struct {
	libConfig.Config
	Port          string `json:"port"`
	DbEngine      string `json:"db_engine"`
	DbConn        string `json:"db_conn"`
	Production    bool   `json:"production"`
	AutoMigration bool   `json:"auto_migration"`
	CorsEnabled   bool   `json:"cors_enabled"`
}

// Configuration is proxy's global configuration
var Configuration *LocalConfig

func init() {
	Configuration = &LocalConfig{}
	// Load the local.ini file
	_, dir, _, _ := runtime.Caller(1)
	if err := Configuration.SetPath(path.Join(path.Dir(dir), "/configs/local.ini")); err != nil {
		log.Fatal(err)
	}

	if err := Configuration.SetPrefix("proxy"); err != nil {
		log.Fatal(err)
	}

	if err := libConfig.LoadConfig(Configuration); err != nil {
		log.Fatal(err)
	}

}
