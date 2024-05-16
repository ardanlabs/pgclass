package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ardanlabs/service/business/sdk/migrate"
	"github.com/ardanlabs/service/business/sdk/sqldb"
)

func main() {
	if err := Migrate(); err != nil {
		log.Fatalln(err)
	}
}

func Migrate() error {
	cfg := sqldb.Config{
		User:         "postgres",
		Password:     "postgres",
		Host:         "database-service.sales-system.svc.cluster.local",
		Name:         "postgres",
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		DisableTLS:   true,
	}

	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")

	if err := migrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
