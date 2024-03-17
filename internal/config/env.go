package config

import (
	"os"
)

const ServerDbName = "SERVER_DB_NAME"
const SuperAdminUsername = "SUPER_ADMIN_USERNAME"
const SuperAdminPassword = "SUPER_ADMIN_PASSWORD"
const JwtSecret = "JWT_SECRET"

// GetEnv Simple helper function to read an environment or return a default value
func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return "ENV_ERROR"
}
