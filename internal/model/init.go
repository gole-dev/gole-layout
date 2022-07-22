package model

import (
	"fmt"
	"sync"

	"github.com/gole-dev/gole/pkg/config"
	"github.com/gole-dev/gole/pkg/storage/orm"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Once sync.Once
)

// Init init db
func Init() *gorm.DB {
	cfg, err := loadConf()
	if err != nil {
		panic(fmt.Sprintf("load orm conf err: %v", err))
	}

	DB = orm.NewMySQL(cfg)
	return DB
}

// GetDB get a db instance
func GetDB() *gorm.DB {
	Once.Do(func() {
		DB = Init()
	})
	return DB
}

// loadConf load gorm config
func loadConf() (ret *orm.Config, err error) {
	var cfg orm.Config
	if err := config.Load("database", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
