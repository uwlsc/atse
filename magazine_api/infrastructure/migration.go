package infrastructure

import (
	"database/sql"
	"fmt"
	"magazine_api/lib"

	_ "github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"
)

//Migrations -> Migration Struct
type Migrations struct {
	logger lib.Logger
	env    lib.Env
	db     Database
}

//NewMigrations -> return new Migrations struct
func NewMigrations(logger lib.Logger, db Database, env lib.Env) Migrations {
	return Migrations{
		logger: logger,
		env:    env,
		db:     db,
	}
}

//Migrate -> migrates all table
func (m Migrations) Migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migration/",
	}

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", m.env.DBUsername,
		m.env.DBPassword, m.env.DBHost, m.env.DBPort, m.env.DBName)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		m.logger.Error("Error in migration", err.Error())
		m.logger.Panic(err)
	}

	defer db.Close()

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		m.logger.Error("Error in migration", err.Error())
	}
}
