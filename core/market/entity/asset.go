package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

func NewAsset(id, name string, marketVol int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVol,
	}

}
