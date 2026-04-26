package migration

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Migration struct {
	Version string
	SQL     string
}

func LoadMigrations(dir string) ([]Migration, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var migrations []Migration

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()

		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}

		version := strings.Split(name, ".")[0]

		path := filepath.Join(dir, name)
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		migrations = append(migrations, Migration{
			Version: version,
			SQL:     string(content),
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}
