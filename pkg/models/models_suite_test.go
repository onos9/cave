package models_test

import (
	"testing"

	"github.com/cave/migrations"
	"github.com/cave/pkg/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = Describe("Configuration", func() {
	AfterSuite(func() {
		models.CloseDB()
	})

	BeforeSuite(func() {
		db := ConnectToTestDatabase()
		migrations.Migrate(db)
	})
})
