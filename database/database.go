package database

import (
	"database/sql"
	"embed"
)

var (
	DbConnection *sql.DB
)

// Embed the migrations directory
//
//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

func DbMigrate(dbParam *sql.DB) {
	// migrations := &migrate.EmbedFileSystemMigrationSource{
	// 	FileSystem: dbMigrations,
	// 	Root:       "sql_migrations",
	// }

	// n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	// if errs != nil {
	// 	log.Fatalf("Failed to apply migrations: %v", errs)
	// }

	DbConnection = dbParam

	//fmt.Println("Applied", n, "migrations!")
}
