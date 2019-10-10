package core

type jsonInputAssetType struct {
	Description string
}

type jsonInputAsset struct {
	TypeID      int64
	Description string
	Geometry    string
	Parent      []int64
}
