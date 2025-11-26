package config

import "time"

const (
	ServerSettingsDefaultHost            = "localhost"
	ServerSettingsDefaultPort            = "8080"
	ServerSettingsDefaultReadTimeout     = 5 * time.Second
	ServerSettingsDefaultWriteTimeout    = 10 * time.Second
	ServerSettingsDefaultShutdownTimeout = 30 * time.Second
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
