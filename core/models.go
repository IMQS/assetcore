package core

import "github.com/IMQS/nf/nfdb"

type assetType struct {
	nfdb.Model
	Description string `gorm:"not null"`
}

type asset struct {
	nfdb.Model
	Description string
	AssetTypeID int64 `gorm:"not null"`
	//AssetType   assetType `gorm:"column:type"`
}
