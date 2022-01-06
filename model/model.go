package model

import (
	"fmt"
	"time"

	"github.com/iojelly/api-server/global"
	"github.com/iojelly/api-server/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 2
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(s,
		databaseSetting.Host,
		databaseSetting.UserName,
		databaseSetting.PassWord,
		databaseSetting.DBName,
		databaseSetting.Port,
		databaseSetting.SSLMode,
		databaseSetting.TimeZone,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(time.Hour).
		SetMaxIdleConns(10).
		SetMaxOpenConns(100))

	return db, nil
}
