package infrastructure

import (
	"context"
	"fmt"
	"magazine_api/lib"

	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	*pgxpool.Pool
}

// Creates a new database instance
func NewDatabase(logger lib.Logger, env lib.Env) Database {
	db, err := initDB(logger, env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	if err != nil {
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	return Database{db}
}

func initDB(logger lib.Logger, username, password, host, port, dbname string) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, dbname)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.Logger = zapadapter.NewLogger(logger.Desugar().Named("postgres"))

	runtimeParams := make(map[string]string)
	runtimeParams["application_name"] = "magazine_api"
	poolConfig.ConnConfig.RuntimeParams = runtimeParams

	poolConfig.AfterConnect = func(c1 context.Context, conn *pgx.Conn) error {
		dt, err := pgxtype.LoadDataType(context.Background(), conn, conn.ConnInfo(), "user_role")
		if err != nil {
			conn.Exec(context.Background(), `CREATE TYPE user_role AS ENUM (
				'user',
				'magazine_manager',
				'employee',
				'accountant',
				'contributor',
				'advertiser',
				'marketing',
				'admin'
			);`)

			dt, err = pgxtype.LoadDataType(context.Background(), conn, conn.ConnInfo(), "user_role")
			if err != nil {
				logger.Info(err)
				return err
			}
		}

		conn.ConnInfo().RegisterDataType(dt)
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (db Database) StopDB() {
	db.Close()
}
