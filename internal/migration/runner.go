package migration

import (
	"database/sql"
	"log"
)

func Run(db *sql.DB, dir string) error {
	migrations, err := LoadMigrations(dir)
	if err != nil {
		return err
	}

	for _, m := range migrations {
		applied, err := isApplied(db, m.Version)
		if err != nil {
			return err
		}

		if applied {
			continue
		}

		log.Println("Applying migration:", m.Version)

		if err := apply(db, m); err != nil {
			return err
		}
	}

	return nil
}

func isApplied(db *sql.DB, version string) (bool, error) {
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM schema_migrations WHERE version = $1
		)
	`, version).Scan(&exists)

	return exists, err
}

func apply(db *sql.DB, m Migration) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(m.SQL)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO schema_migrations(version)
		VALUES ($1)
	`, m.Version)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
