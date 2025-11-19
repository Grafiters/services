package lib

import (
	"fmt"
	"log"
	"os"
	env "riskmanagement/lib/env"
	"time"

	zapLog "gitlab.com/golang-package-library/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database modal
type Database struct {
	DB *gorm.DB
	// DBRaw *sql.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env env.Env, zapLogger zapLog.Logger) Database {

	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		zapLogger.Zap.Info("Url: ", url)
		LogChecklist("Mysql Connection Refused", false)
		zapLogger.Zap.Panic(err)
	}

	// zapLogger.Zap.Info("Database connection established")
	LogChecklist("Mysql Connection Established", true)

	return Database{
		DB: db,
	}
}
