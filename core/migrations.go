package core

import (
	"github.com/IMQS/log"
	"github.com/IMQS/nf"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var migrationsSQL []string

func openDB(log *log.Logger) (*gorm.DB, error) {
	driver := "postgres"
	dsn := "user=unit_test_user password=unit_test_password dbname=testgorm sslmode=disable"
	var flags nf.DBConnectFlags
	//flags |= nf.DBConnectFlagWipeDB
	db, err := nf.OpenDB(log, driver, dsn, nf.MakeMigrations(log, migrationsSQL), flags)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Building{}, &Floor{})
	return db, nil
}

func init() {
	migrationsSQL = []string{
		`
		-- CREATE TABLE buildings (id BIGSERIAL PRIMARY KEY, name VARCHAR);
		-- CREATE TABLE floors (id BIGSERIAL PRIMARY KEY, building_id BIGINT NOT NULL, floor_name VARCHAR);
		`,
	}
}
