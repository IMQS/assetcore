package core

import (
	"github.com/IMQS/log"
	"github.com/IMQS/nf/nfdb"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// If you modify this schema, keep it up to date at https://dbdiagram.io/d/5d9f86ffff5115114db5209a

var migrationsSQL []string

func openDB(log *log.Logger, dbConf nfdb.DBConfig) (*gorm.DB, error) {
	var flags nfdb.DBConnectFlags
	flags |= nfdb.DBConnectFlagWipeDB
	db, err := nfdb.OpenDB(log, dbConf.Driver, dbConf.DSN(), nfdb.MakeMigrations(log, migrationsSQL), flags)
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

		CREATE TABLE asset (id BIGSERIAL PRIMARY KEY, description VARCHAR, geom geometry(GeometryZ, 4326), asset_type_id BIGINT);
		CREATE INDEX idx_asset_geom ON asset USING GIST (geom);

		CREATE TABLE asset_type (id BIGSERIAL PRIMARY KEY, description VARCHAR);
		CREATE UNIQUE INDEX idx_asset_type_description ON asset_type (description);
		
		CREATE TABLE hierarchy1 (id BIGINT PRIMARY KEY, parent_id BIGINT);
		CREATE INDEX idx_hierarchy1_parent ON hierarchy1(parent_id);
		`,
	}
}
