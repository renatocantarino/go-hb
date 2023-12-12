package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	AssertId string
	Shares   int
}

func NewInvestor(id, name string) *Investor {

	return &Investor{
		ID:            id,
		Name:          name,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func NewAssetPosition(assetId string, share int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssertId: assetId,
		Shares:   share,
	}

}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetId string, share int) {
	assets := i.GetAssetPosition(assetId)
	if assets == nil {
		i.AssetPosition = append(i.AssetPosition, NewAssetPosition(assetId, share))
	} else {
		assets.Shares += share
	}
}

func (i *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {

	for _, item := range i.AssetPosition {
		if item.AssertId == assetId {
			return item
		}
	}
	return nil
}
