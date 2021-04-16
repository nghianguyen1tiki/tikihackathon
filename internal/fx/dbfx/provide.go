package dbfx

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/pkg/assert"
	"github.com/nghiant3223/tikihackathon/pkg/log"
)

func provideDB() (*gorm.DB, error) {
	dsn := viper.GetString("db.dsn")
	assert.NotEmpty(dsn, "db.dsn is empty")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorw("cannot open database connection", "error", err)
		return nil, err
	}
	return db, nil
}
