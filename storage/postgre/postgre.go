package postgre

import (
	"archi/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(conf *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(conf.DB.DSN), &gorm.Config{})
}
