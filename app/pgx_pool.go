package app

import (
	"context"
	"fmt"
	"os"
	"simple-attendance/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(cfg *config.Config) *pgxpool.Pool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s`,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("success to connect database pgxpool")
	}

	return dbpool
}
