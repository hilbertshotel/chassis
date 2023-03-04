package dep

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// PLACE ALL DEPENDENCIES IN THIS FUNCTION
func Init() (*Dependencies, error) {

	// 1.
	log := initLogger()

	// 2.
	cfg, err := loadConfig("config.json")
	if err != nil {
		return nil, err
	}

	// 3.
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		return nil, err
	}

	// 4.
	tmp, err := template.ParseGlob(cfg.TmpDir)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Log: log,
		Cfg: cfg,
		DB:  db,
		Tmp: tmp,
	}, nil
}

// INITIALIZE LOGGER
func initLogger() *Logger {
	okLog := log.New(os.Stdout, "OK ", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)

	return &Logger{
		Error: func(err error) {
			_, file, line, _ := runtime.Caller(1)
			msg := fmt.Sprintf("(%v %v) %v", filepath.Base(file), line, err)
			errLog.Println(msg)
		},

		Ok: func(msg string) {
			okLog.Println(msg)
		},
	}
}

// LOAD CONFIG
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
