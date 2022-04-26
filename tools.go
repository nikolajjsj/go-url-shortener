//go:build tools
// +build tools

package tools

import (
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/google/uuid"
	_ "github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)
