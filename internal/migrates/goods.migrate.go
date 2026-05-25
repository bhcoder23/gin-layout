package migrates

import (
	"github.com/bhcoder23/gin-layout/internal/models"
	"gorm.io/gorm"
)

func init() {
	regMigrate(func(d *gorm.DB) error {
		return d.AutoMigrate(&models.Goods{})
	})
}
