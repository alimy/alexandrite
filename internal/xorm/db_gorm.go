// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package xorm

import (
	"fmt"
	"strings"
	"time"

	"github.com/alimy/hori/dao"
	"github.com/alimy/hori/internal/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type dbr struct {
	*gorm.DB
}

func (d *dbr) Whoami() string {
	return "gorm"
}

func initDB() (dao.Repository, error) {
	config := conf.MyConfig()

	level := logger.Info
	if conf.IsProdMode() {
		level = logger.Warn
	}

	// NOTE: AutoMigrate does not respect logger passed in gorm.Config.
	logger.Default = logger.New(logrus.New(), logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      level,
	})

	db, err := openDB(config.Database, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Microsecond)
		},
	})
	if err != nil {
		return nil, fmt.Errorf("open database failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying *sql.DB failed: %w", err)
	}
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	switch config.Database.Type {
	case "postgres":
		conf.UsePostgreSQL = true
	case "mysql":
		conf.UseMySQL = true
		db = db.Set("gorm:table_options", "ENGINE=InnoDB").Session(&gorm.Session{})
	case "sqlite3":
		conf.UseSQLite3 = true
	case "mssql":
		conf.UseMSSQL = true
	default:
		panic("unreachable")
	}

	// NOTE: GORM has problem detecting existing columns, see https://github.com/gogs/gogs/issues/6091.
	// Therefore only use it to create new tables, and do customized migration with future changes.
	for _, table := range dao.Tables() {
		if db.Migrator().HasTable(table) {
			continue
		}
		name := strings.TrimPrefix(fmt.Sprintf("%T", table), "*db.")
		err = db.Migrator().AutoMigrate(table)
		if err != nil {
			return nil, fmt.Errorf("auto migrate %q failed: %w", name, err)
		}
		logrus.Trace("Auto migrated %q", name)
	}

	return &dbr{
		DB: db,
	}, nil
}

func openDB(opts conf.Database, cfg *gorm.Config) (*gorm.DB, error) {
	dsn, err := opts.Dsn()
	if err != nil {
		return nil, fmt.Errorf("parse DSN failed %w", err)
	}

	var dialector gorm.Dialector
	switch opts.Type {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "mssql":
		dialector = sqlserver.Open(dsn)
	case "sqlite3":
		dialector = sqlite.Open(dsn)
	default:
		panic("unreachable")
	}

	return gorm.Open(dialector, cfg)
}
