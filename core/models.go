package core

import "github.com/IMQS/nf"

type assetType struct {
	nf.Model
	Description string
}

type asset struct {
	nf.Model
	Description string
	AssetType   assetType `gorm:"column:type"`
}
