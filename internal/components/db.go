package components

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := viper.GetString("mysql.dsn")
	if dsn == "" {
		panic("please add mysql.dsn!")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("db connect error: ", zap.Error(err))
		return
	}
	zap.L().Info("db connect success: ", zap.String("dsn", dsn))
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	DB = db.Debug()
}
