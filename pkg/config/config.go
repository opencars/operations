package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

const DefaultInsertBatchSize = 10000

// Settings is decoded configuration file.
type Settings struct {
	DB     Database `yaml:"database"`
	Worker Worker   `yaml:"worker"`
	Log    Log      `yaml:"log"`
	Server Server   `yaml:"server"`
}

// Log represents settings for application logger.
type Log struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

// Server represents settings for creating http server.
type Server struct {
	ShutdownTimeout Duration `yaml:"shutdown_timeout"`
	ReadTimeout     Duration `yaml:"read_timeout"`
	WriteTimeout    Duration `yaml:"write_timeout"`
	IdleTimeout     Duration `yaml:"idle_timeout"`
}

// Database contains configuration details for database.
type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

// Worker contains settings for data processing by cmd/worker.
type Worker struct {
	PackageID       string `yaml:"package_id"`
	InsertBatchSize int    `yaml:"insert_batch_size"`
}

// Address return API address in "host:port" format.
func (db *Database) Address() string {
	return db.Host + ":" + strconv.Itoa(db.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	var config Settings

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	if config.Worker.InsertBatchSize == 0 {
		config.Worker.InsertBatchSize = DefaultInsertBatchSize
	}

	return &config, nil
}
