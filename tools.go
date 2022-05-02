//go:build tools
// +build tools

package tools

import (
	_ "github.com/gofiber/fiber/v2"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)
