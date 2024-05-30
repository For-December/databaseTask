package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
)

var EnvCfg = struct {
	LogLevel      string `env:"LOG_LEVEL" envDefault:"debug"`
	MysqlHost     string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	MysqlPort     string `env:"MYSQL_PORT" envDefault:"3306"`
	MysqlUser     string `env:"MYSQL_USER" envDefault:"root"`
	MysqlPassword string `env:"MYSQL_PASSWORD"`
	MysqlDatabase string `env:"MYSQL_DATABASE"`
	AutoMigrate   bool   `env:"AUTO_MIGRATE"`
}{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can not read env from file system, please check the right this program owned.")
	}

	//EnvCfg = envConfig{}

	if err := env.Parse(&EnvCfg); err != nil {
		panic("Can not parse env from file system, please check the env.")
	}

}
