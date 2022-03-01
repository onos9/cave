package mods

import (
	"github.com/jinzhu/gorm"
)

var (
	//errHandlerNotSet error = errors.New("handler not set properly")
	handler *gorm.DB
)

// SetRepoDB global db handler
func SetRepoDB(db *gorm.DB) {
	handler = db
}

// CloseDB closes handler db
func CloseDB() {
	if handler != nil {
		handler.Close()
	}
}
