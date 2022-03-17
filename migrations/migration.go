package migrations

import (
	"github.com/cave/cmd/api/mods"
	"github.com/cave/pkg/models"

	"github.com/jinzhu/gorm"
)

// Migrate migrates gorm models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.AnswerOption{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Certificate{})
	db.AutoMigrate(&models.IssuedCertificate{})
	db.AutoMigrate(&models.ContentBlock{})
	db.AutoMigrate(&models.CourseAuthor{})
	db.AutoMigrate(&models.Course{})
	db.AutoMigrate(&models.EvaluationCriteria{})
	db.AutoMigrate(&models.Level{})
	db.AutoMigrate(&models.QuizQuestion{})
	db.AutoMigrate(&models.QuizUserAnswer{})
	db.AutoMigrate(&models.Quiz{})
	db.AutoMigrate(&models.Target{})
	db.AutoMigrate(&models.TargetVersion{})
	db.AutoMigrate(&models.TargetVersion{})
	db.AutoMigrate(&models.TargetGroup{})
	db.AutoMigrate(&models.StudentCourse{})

	db.AutoMigrate(&mods.Category{})
	db.AutoMigrate(&mods.Channel{})
	db.AutoMigrate(&mods.Comment{})
	db.AutoMigrate(&mods.Dislike{})
	db.AutoMigrate(&mods.Like{})
	db.AutoMigrate(&mods.User{})
	db.AutoMigrate(&mods.Subscription{})
	db.AutoMigrate(&mods.Video{})
	db.AutoMigrate(&mods.View{})

	db.AutoMigrate(&mods.Background{})
	db.AutoMigrate(&mods.Bio{})
	db.AutoMigrate(&mods.Candidate{})
	db.AutoMigrate(&mods.Qualification{})
	db.AutoMigrate(&mods.Terms{})
	db.AutoMigrate(&mods.Ref{})
	db.AutoMigrate(&mods.Health{})
}
