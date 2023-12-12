package entity

type Asset struct {
	Id           string
	Name         string
	MarketVolume int
}

func NewAsset(id, name string, marketVol int) *Asset {
	return &Asset{
		Id:           id,
		Name:         name,
		MarketVolume: marketVol,
	}

}
