package ay

import (
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {
	if err := GetDB(); err != nil {
		Logger.Error(err.Error())
	}
}

func Begin() *gorm.DB {
	return Db.Begin()
}

func GetDB() (err error) {
	InitializeRedis()
	var option gorm.Dialector

	switch Yaml.GetString("sql.type") {
	case "mysql":
		dsn := Yaml.GetString("sql.user") + ":" + Yaml.GetString("sql.password") + "@tcp(" + Yaml.GetString("sql.localhost") + ":" + Yaml.GetString("sql.port") + ")/" + Yaml.GetString("sql.database") + "?charset=utf8mb4&parseTime=true&loc=Local"
		option = mysql.New(mysql.Config{DSN: dsn})
	case "sqlite":
		option = sqlite.Open(Yaml.GetString("sql.localhost"))
	default:

	}

	if Db, err = gorm.Open(option, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default,
	}); err != nil {
		return
	}

	Db.Logger = NewGormLogger()

	return
}

func IgnoreNotFound(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if errors.Is(err, redis.Nil) {
		return nil
	}
	return err
}

func IgnoreNotFoundReturn[T any](r T, err error) (T, error) {
	err = IgnoreNotFound(err)
	return r, err
}
