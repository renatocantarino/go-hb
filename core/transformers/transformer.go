package transformers

import (
	"github.com/renatocantarino/go-home-broker/core/dto"
	"github.com/renatocantarino/go-home-broker/core/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {

	asset := entity.NewAsset(input.AssetId, input.AssetId, 1000)
	investor := entity.NewInvestor(input.InvestorId, input.InvestorId)
	order := entity.NewOrder(input.OrderId, investor, asset, input.Share, input.Price, input.OrderType)

	if input.CurrentShare > 0 {
		position := entity.NewAssetPosition(asset.Id, input.CurrentShare)
		investor.AddAssetPosition(position)
	}

	return order

}
