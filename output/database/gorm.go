package database

import (
	"errors"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBClickhouse = "clickhouse"
	DBPostgres   = "postgres"
	DBMySQL      = "mysql"
)

func NewGormDB(dbType string, dsn string) (*gorm.DB, error) {
	switch dbType {
	case DBClickhouse: // dsn := "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20"
		return getGormDB(clickhouse.Open(dsn))
	case DBPostgres: // dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		return getGormDB(postgres.Open(dsn))
	case DBMySQL: // dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		return getGormDB(mysql.Open(dsn))
	default:
		return nil, errors.New("unsupported database type")
	}
}

func getGormDB(d gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(d, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
