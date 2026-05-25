package components

import (
	"errors"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := viper.GetString("mysql.dsn")
	if dsn == "" {
		return errors.New("missing mysql.dsn")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	zap.L().Info("db connected", zap.String("dsn", RedactDSN(dsn)))
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db.Debug()
	return nil
}

// RedactDSN removes credentials from DSNs before they are written to logs.
func RedactDSN(dsn string) string {
	at := strings.Index(dsn, "@")
	if at < 0 {
		return dsn
	}
	prefix := dsn[:at]
	colon := strings.LastIndex(prefix, ":")
	if colon < 0 {
		return dsn
	}
	return prefix[:colon+1] + "***" + dsn[at:]
}
