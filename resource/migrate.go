package resource

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// ponytail: up-only runner; .down.sql files are kept for reference but rollback
// is not wired. Add a Rollback() + CLI command when you need to undo.
func Migrate(db *gorm.DB) error {
	if err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (version BIGINT PRIMARY KEY, applied_at TIMESTAMPTZ NOT NULL DEFAULT now())`).Error; err != nil {
		return err
	}

	var current int64
	if err := db.Raw(`SELECT COALESCE(MAX(version), 0) FROM schema_migrations`).Scan(&current).Error; err != nil {
		return err
	}

	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return err
	}

	versions := make([]int, 0, len(entries))
	up := map[int]string{}
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		version, err := strconv.Atoi(strings.SplitN(name, "_", 2)[0])
		if err != nil {
			return fmt.Errorf("parse migration %s: %w", name, err)
		}
		content, err := migrationsFS.ReadFile("migrations/" + name)
		if err != nil {
			return err
		}
		up[version] = string(content)
		versions = append(versions, version)
	}
	sort.Ints(versions)

	for _, v := range versions {
		if v <= int(current) {
			continue
		}
		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(up[v]).Error; err != nil {
				return err
			}
			return tx.Exec(`INSERT INTO schema_migrations (version) VALUES (?)`, v).Error
		}); err != nil {
			return fmt.Errorf("migration %d: %w", v, err)
		}
	}
	return nil
}
