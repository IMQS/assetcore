package core

import (
	"github.com/IMQS/log"
	"github.com/IMQS/nf"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var migrationsSQL []string

func openDB(log *log.Logger, dbConf nf.DBConfig) (*gorm.DB, error) {
	var flags nf.DBConnectFlags
	flags |= nf.DBConnectFlagWipeDB
	db, err := nf.OpenDB(log, dbConf.Driver, dbConf.DSN(), nf.MakeMigrations(log, migrationsSQL), flags)
	if err != nil {
		return nil, err
	}
	//db.AutoMigrate(&Building{}, &Floor{})
	return db, nil
}

func init() {
	migrationsSQL = []string{
		`
		CREATE EXTENSION IF NOT EXISTS postgis;

		CREATE TABLE asset (id BIGSERIAL PRIMARY KEY, description VARCHAR, geom geometry(GeometryZ, 4326), type_id BIGINT);
		CREATE INDEX idx_asset_geom ON asset USING GIST (geom);

		CREATE TABLE type (id BIGSERIAL PRIMARY KEY, description VARCHAR);
		
		CREATE TABLE hierarchy1 (id BIGINT PRIMARY KEY, parent_id BIGINT);
		CREATE INDEX idx_hierarchy1_parent ON hierarchy1(parent_id);
		`,
	}
}
