package database

import (
	"fmt"
	"log"

	"github.com/cave/configs"
	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

type Database struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Initialize gets the config and returns a database pointer
func Initialize(conf configs.Storage) (*Database, error) {
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.Dbuser, conf.Dbpassword, conf.Database)

	log.Println("main : Initialize Redis")
	redisClient := redis.NewClient(&redis.Options{
		Addr:        configs.CFG.Redis.Host,
		DB:          configs.CFG.Redis.DB,
		DialTimeout: configs.CFG.Redis.DialTimeout,
	})

	db, err := gorm.Open("postgres", url)
	return &Database{DB: db, Redis: redisClient}, err
}

// InjectDB injects database to gin server
func InjectDB(db *Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
