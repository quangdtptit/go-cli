package migration

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func New(path string) (*migrate.Migrate, error) {
	URL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "123456aA@", "localhost", 3542, "test")
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
