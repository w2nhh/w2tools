package ormsql

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	MYSQLDSN = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&autocommit=true"
)

type SqlOptions struct {
	Address  string
	User     string
	Password string
	Db       string
	// LifeTime time.Duration
}

type SqlOption func(opt *SqlOptions)

func SqlAddrOptions(addr string) SqlOption {
	return func(opt *SqlOptions) {
		opt.Address = addr
	}
}
func SqlUserOptions(user string) SqlOption {
	return func(opt *SqlOptions) {
		opt.User = user
	}
}
func SqlPWDOptions(password string) SqlOption {
	return func(opt *SqlOptions) {
		opt.Password = password
	}
}
func SqlDBOptions(dbn string) SqlOption {
	return func(opt *SqlOptions) {
		opt.Db = dbn
	}
}

func NewDBCli(cfg *SqlOptions) *gorm.DB {
	cli, err := gorm.Open("mysql", fmt.Sprintf(MYSQLDSN, cfg.User, cfg.Password, cfg.Address, cfg.Db))
	if err != nil {
		panic(err.Error())
	}

	return cli
}
