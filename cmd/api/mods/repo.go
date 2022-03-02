package mods

import (
	"github.com/cave/pkg/database"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	//errHandlerNotSet error = errors.New("handler not set properly")
	handler *gorm.DB
	RedisClient *redis.Client
)

// SetRepoDB global db handler
func SetRepoDB(db *database.Database) {
	handler = db.DB
	RedisClient = db.Redis
}

// CloseDB closes handler db
func CloseDB() {
	if handler != nil {
		handler.Close()
	}
}
