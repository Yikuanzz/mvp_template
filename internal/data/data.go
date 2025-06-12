package data

import (
	"fmt"

	"mvp/config"
	"mvp/internal/handler"
	"mvp/utils/log"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewUserRepo,
)

type Data struct {
	db  *gorm.DB
	log log.Logger
}

// NewUserRepo 创建用户仓库
func NewUserRepo(data *Data, log log.Logger) handler.UserRepo {
	return &UserRepo{data: data, log: log}
}

// NewData 创建数据
func NewData(db *gorm.DB, log log.Logger) *Data {
	return &Data{db: db, log: log}
}

// NewDB 创建数据库连接
func NewDB(conf *config.Config, log log.Logger) *gorm.DB {
	log.Info("connecting to database", "host", conf.MySQL.Host, "port", conf.MySQL.Port, "database", conf.MySQL.Database)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.MySQL.Username, conf.MySQL.Password, conf.MySQL.Host, conf.MySQL.Port, conf.MySQL.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database", "error", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Error("failed to migrate database", "error", err)
	}

	return db
}
