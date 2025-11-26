package config

import (
	"fmt"
	"time"
)

const (
	ServerSettingsDefaultHost            = "localhost"
	ServerSettingsDefaultPort            = "8080"
	ServerSettingsDefaultReadTimeout     = 5 * time.Second
	ServerSettingsDefaultWriteTimeout    = 10 * time.Second
	ServerSettingsDefaultShutdownTimeout = 30 * time.Second

	DatabaseSettingsDefaultHost            = "localhost"
	DatabaseSettingsDefaultPort            = 5432
	DatabaseSettingsDefaultUser            = "postgres"
	DatabaseSettingsDefaultPassword        = "postgres"
	DatabaseSettingsDefaultDBName          = "atlans"
	DatabaseSettingsDefaultSSLMode         = "disable"
	DatabaseSettingsDefaultMinConns        = 10
	DatabaseSettingsDefaultMaxConns        = 100
	DatabaseSettingsDefaultConnMaxLifetime = 5 * time.Minute
	DatabaseSettingsDefaultConnMaxIdleTime = 1 * time.Minute
	DatabaseSettingsDefaultMigrationsPath  = "file://migrations"
)

type ServerSettings struct {
	Host            string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

func (s *ServerSettings) SetDefaults() {
	if s.Host == "" {
		s.Host = ServerSettingsDefaultHost
	}

	if s.Port == "" {
		s.Port = ServerSettingsDefaultPort
	}

	if s.ReadTimeout == 0 {
		s.ReadTimeout = ServerSettingsDefaultReadTimeout
	}

	if s.WriteTimeout == 0 {
		s.WriteTimeout = ServerSettingsDefaultWriteTimeout
	}

	if s.ShutdownTimeout == 0 {
		s.ShutdownTimeout = ServerSettingsDefaultShutdownTimeout
	}
}

type DatabaseSettings struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MinConns        int
	MaxConns        int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MigrationsPath  string
}

func (d *DatabaseSettings) SetDefaults() {
	if d.Host == "" {
		d.Host = DatabaseSettingsDefaultHost
	}

	if d.Port == 0 {
		d.Port = DatabaseSettingsDefaultPort
	}

	if d.User == "" {
		d.User = DatabaseSettingsDefaultUser
	}

	if d.Password == "" {
		d.Password = DatabaseSettingsDefaultPassword
	}

	if d.DBName == "" {
		d.DBName = DatabaseSettingsDefaultDBName
	}

	if d.SSLMode == "" {
		d.SSLMode = DatabaseSettingsDefaultSSLMode
	}

	if d.MinConns == 0 {
		d.MinConns = DatabaseSettingsDefaultMinConns
	}

	if d.MaxConns == 0 {
		d.MaxConns = DatabaseSettingsDefaultMaxConns
	}

	if d.ConnMaxLifetime == 0 {
		d.ConnMaxLifetime = DatabaseSettingsDefaultConnMaxLifetime
	}

	if d.ConnMaxIdleTime == 0 {
		d.ConnMaxIdleTime = DatabaseSettingsDefaultConnMaxIdleTime
	}

	if d.MigrationsPath == "" {
		d.MigrationsPath = DatabaseSettingsDefaultMigrationsPath
	}
}

func (d *DatabaseSettings) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode)
}
