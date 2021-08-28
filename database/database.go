package database

import (
	"fmt"
	"github.com/aURORA-JC/headerfaker/model"
	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// InitDataBase Must be use at function main() to init the database connection
func InitDataBase(cfg *model.MySQLConfig) (dbError error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbError != nil {
		color.Red("ERROR! Failed to initialize database! exited")
		os.Exit(001)
	}
	return
}
