package dep

import (
	"database/sql"
	"html/template"
)

// LOGGER
type Logger struct {
	Error func(error)
	Ok    func(string)
}

// CONFIG
type Config struct {
	ConnStr  string
	HostAddr string
	TmpDir   string
}

// DEPENDENCIES
type Dependencies struct {
	Log *Logger
	Cfg *Config
	DB  *sql.DB
	Tmp *template.Template
}
