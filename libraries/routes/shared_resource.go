package routes

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

// SharedResource stores the shared data that needs to pass into the controller
type SharedResource struct {
	DB          *gorm.DB
	Validator   *validator.Validate
	RedisClient *redis.Client
}
