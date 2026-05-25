// Package migrates manages database schema migrations.
package migrates

import (
	"fmt"

	"github.com/bhcoder23/gin-layout/internal/components"
	"gorm.io/gorm"
)

type migrate func(*gorm.DB) error

var routers = []migrate{}

func DoMigrate() error {
	for index, route := range routers {
		if err := route(components.DB); err != nil {
			return fmt.Errorf("run migration %d: %w", index, err)
		}
	}
	return nil
}

func regMigrate(r ...migrate) {
	routers = append(routers, r...)
}
