package app

import (
	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/gormadp"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBA2 *gormadp.DBAdapter

// Start database connection

func db2Up(args3 *AppArgs) {
	var loglevel logger.LogLevel
	if config.Of.App.Debug() {
		loglevel = logger.Info
	} else {
		loglevel = logger.Silent
	}

	pkgOptions := &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	}

	cfg := gormadp.Config{
		Host:     config.Of.Database2.Host,
		Port:     config.Of.Database2.Port,
		User:     config.Of.Database2.User,
		Password: config.Of.Database2.Password,
		DBName:   config.Of.Database2.Name,
		Dialect:  gormadp.Dialect(config.Of.Database2.Dialect),
		Options:  pkgOptions,
	}

	opts := cfg.Dialect.PgOptions(gormadp.PgConfig{
		SSLMode:  false,
		TimeZone: "Asia/Manila",
	})

	dba := gormadp.Open(cfg, opts)

	DBA2 = dba
	printOutUp("New DB 2 connection successfully open")
}

// Stop database connection
func db2Down() {
	printOutDown("Closing current DB 2 connection...")
	DBA2.Close()
}
