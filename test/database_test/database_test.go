package database_test

import (
	"mongosteen/internal/database"
	"testing"
)

func BenchmarkCrud(b *testing.B) {
	database.Connect()
	defer database.Close()
	// database.CreateMigration()
	database.Migrate()
	for i := 0; i < b.N; i++ {
		database.Crud()
	}
}
