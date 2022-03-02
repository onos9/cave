package models

import (
	"errors"

	"github.com/cave/pkg/database"
	"github.com/jinzhu/gorm"
)

var (
	errHandlerNotSet error = errors.New("handler not set properly")
	handler          *gorm.DB
)

// SetRepoDB global db handler
func SetRepoDB(db *database.Database) {
	handler = db.DB
}

// CloseDB closes handler db
func CloseDB() {
	if handler != nil {
		handler.Close()
	}
}
