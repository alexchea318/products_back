package config

import (
	"os"
)

const ServerDbName = "SERVER_DB_NAME"
const DbClient = "DB_CLIENT"
const SuperAdminUsername = "SUPER_ADMIN_USERNAME"
const SuperAdminPassword = "SUPER_ADMIN_PASSWORD"
const JwtSecret = "JWT_SECRET"
const AuthorizationHeader = "AUTH_HEADER"
const GinMode = "GIN_MODE"
const GinPort = "PORT"

const ReleaseMode = "release"

// GetEnv Simple helper function to read an environment or return a default value
func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return "ENV_ERROR"
}
