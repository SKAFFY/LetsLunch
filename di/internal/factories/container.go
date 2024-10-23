package factories

import (
	"cmd/di/internal/lookup"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func CreateLogger(ctx context.Context, c lookup.Container) *log.Logger {
	return log.Default()
}

func CreateDB(ctx context.Context, c lookup.Container) *sql.DB {
	psqlInfo := c.Config(ctx).DatabaseURL
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
