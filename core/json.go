package core

type jsonInputAssetType struct {
	Description string
}

type jsonInputAsset struct {
	AssetTypeID int64
	Description string
	Geometry    string
	Parent      []int64
}
