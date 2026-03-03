package migration

import (
	"errors"
	"fmt"

	"github.com/quangdtptit/go-cli/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func New(path string) (*migrate.Migrate, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("error init config: %w", err)
	}

	URL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database)
	m, err := migrate.New("file://"+path, URL)

	if err != nil {
		return nil, fmt.Errorf("error creating migrate instance: %w", err)
	}
	return m, nil
}

func Up(path string) error {
	m, err := New(path)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {

		return fmt.Errorf("error applying up migrations: %w", err)
	}

	return nil
}

func Down(path string) error {
	m, err := New(path)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {

		return fmt.Errorf("error applying down migrations: %w", err)
	}

	return nil
}
