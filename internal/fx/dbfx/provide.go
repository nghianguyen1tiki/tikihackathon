package dbfx

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/config"

	"github.com/nghiant3223/tikihackathon/pkg/log"
)

func provideDB() (*gorm.DB, error) {
	var cfg config.DB
	err := viper.UnmarshalKey("db", &cfg)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Errorw("cannot open database connection", "error", err)
		return nil, err
	}
	return db, nil
}
