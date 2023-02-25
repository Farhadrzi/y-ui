package database

import (
	"io/fs"
	"log"
	"os"
	"path"
	"time"
	"x-ui/database/model"
	"x-ui/xray"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initUser() error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	var count int64
	err = db.Model(&model.User{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		user := &model.User{
			Username: "admin",
			Password: "admin",
		}
		return db.Create(user).Error
	}
	return nil
}

func initInbound() error {
	return db.AutoMigrate(&model.Inbound{})
}

func initSetting() error {
	return db.AutoMigrate(&model.Setting{})
}
func initInboundClientIps() error {
	return db.AutoMigrate(&model.InboundClientIps{})
}
func initClientTraffic() error {
	return db.AutoMigrate(&xray.ClientTraffic{})
}

func initMirrors() error {
	return db.AutoMigrate(&model.Mirror{})
}

func InitDB(dbPath string) error {

	dir := path.Dir(dbPath)
	err := os.MkdirAll(dir, fs.ModeDir)
	if err != nil {
		return err
	}

	dsn := "postgresql://masteradminking:wxpMSyWNnpTCKymDruablA@black-sponge-7079.8nj.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	err = initUser()
	if err != nil {
		return err
	}
	err = initInbound()
	if err != nil {
		return err
	}
	err = initSetting()
	if err != nil {
		return err
	}
	err = initInboundClientIps()
	if err != nil {
		return err
	}
	err = initClientTraffic()
	if err != nil {
		return err
	}

	err = initMirrors()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
