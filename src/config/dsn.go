package config

import (
	"fmt"
	"strings"
)

var dsnConfigs = []string{
	fmt.Sprintf("host=%s", GetEnv("POSTGRES_HOST")),
	fmt.Sprintf("user=%s", GetEnv("POSTGRES_USER")),
	fmt.Sprintf("database=%s", GetEnv("POSTGRES_DB")),
	fmt.Sprintf("password=%s", GetEnv("POSTGRES_PASSWORD")),
	fmt.Sprintf("port=%s", GetEnv("POSTGRES_PORT")),
	"sslmode=disable",
	fmt.Sprintf("TimeZone=%s", GetEnv("POSTGRES_TZ")),
}

var DSN string = strings.Join(dsnConfigs, " ")
