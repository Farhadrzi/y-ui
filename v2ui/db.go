package v2ui

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var v2db *gorm.DB

func initDB(dbPath string) error {
	var err error
	dsn := "postgresql://masteradminking:wxpMSyWNnpTCKymDruablA@black-sponge-7079.8nj.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	v2db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func getV2Inbounds() ([]*V2Inbound, error) {
	inbounds := make([]*V2Inbound, 0)
	err := v2db.Model(V2Inbound{}).Find(&inbounds).Error
	return inbounds, err
}
