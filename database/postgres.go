package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *pgxpool.Pool
var gdb *gorm.DB

type Postgres struct {
	DB  *pgxpool.Pool
	GDB *gorm.DB
}

func NewPostgres() DB {
	return &Postgres{}
}

func (p *Postgres) Set() error {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	config, err := p.SetConfig(connString)
	if err != nil {
		return fmt.Errorf("unable to parse config: %v", err)
	}

	db, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	if err := p.Gorm(db); err != nil {
		return fmt.Errorf("unable to set up gorm db: %v", err)
	}

	return nil
}

func (p *Postgres) Get() *Postgres {
	return &Postgres{
		DB:  db,
		GDB: gdb,
	}
}

func (p *Postgres) Close() error {
	return nil
}

func (p *Postgres) Ping() error {
	return nil
}

func (p *Postgres) Gorm(db *pgxpool.Pool) error {

	var err error

	pg := stdlib.OpenDB(*db.Config().ConnConfig)
	if gdb, err = gorm.Open(postgres.New(postgres.Config{
		Conn: pg,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	}); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) SetConfig(conStr string) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(conStr)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 50
	config.MinConns = 10
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 15 * time.Minute
	config.HealthCheckPeriod = time.Minute

	return config, nil
}
