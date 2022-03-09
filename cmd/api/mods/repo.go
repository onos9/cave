package mods

import (
	"github.com/cave/pkg/database"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	//errHandlerNotSet error = errors.New("handler not set properly")
	handler     *gorm.DB
	RedisClient *redis.Client
)

var categories = []Category{
	{Title: "Systematic Theology I", Semester: 1},
	{Title: "Systematic Theology II", Semester: 1},
	{Title: "Old Testament Survey", Semester: 1},
	{Title: "New Testament Survey", Semester: 1},
	{Title: "Church History Survey", Semester: 1},
	{Title: "Hermeneutics", Semester: 1},
	{Title: "HOmiletics", Semester: 1},
	{Title: "Research & Writing", Semester: 1},
	{Title: "Apostolic Discipleship & Mentoring", Semester: 1},
	{Title: "Introduction To Philosophy", Semester: 1},
	{Title: "Christian Ethics", Semester: 2},
	{Title: "Christian Family", Semester: 2},
	{Title: "Christian Apologetics", Semester: 2},
	{Title: "Introduction To Islam", Semester: 2},
	{Title: "Cross Cultural Missions", Semester: 2},
	{Title: "Principles Of Counseling", Semester: 2},
	{Title: "Children Ministry", Semester: 2},
	{Title: "Youth Ministry", Semester: 2},
	{Title: "Leadership & Administration", Semester: 2},
	{Title: "Africa Traditional Religion (ATR) & World Religions", Semester: 2},
}

// SetRepoDB global db handler
func SetRepoDB(db *database.Database) {
	handler = db.DB
	RedisClient = db.Redis

	var category Category
	for _, cat := range categories {
		category.Title = cat.Title
		category.Semester = cat.Semester
		category.Create()
	}
}

// CloseDB closes handler db
func CloseDB() {
	if handler != nil {
		handler.Close()
	}
}
