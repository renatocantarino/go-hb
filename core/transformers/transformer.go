package transformers

import (
	"github.com/renatocantarino/go-home-broker/core/dto"
	"github.com/renatocantarino/go-home-broker/core/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {

	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID, input.InvestorID)
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		position := entity.NewAssetPosition(asset.ID, input.CurrentShares)
		investor.AddAssetPosition(position)
	}

	return order

}

func TransformOutput(order *entity.Order) *dto.OrderOutput {

	output := &dto.OrderOutput{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  order.OrderType,
		Status:     order.Status,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput

	for _, item := range order.Transactions {

		transactions := &dto.TransactionOutput{
			TransactionID: item.ID,
			BuyerID:       item.BuyingOrder.Investor.ID,
			SellerID:      item.SellingOrder.Investor.ID,
			AssetID:       item.SellingOrder.Asset.ID,
			Price:         item.Price,
			Shares:        item.SellingOrder.Shares - item.SellingOrder.PendingShares,
		}

		transactionsOutput = append(transactionsOutput, transactions)

	}

	output.TransactionsOutput = transactionsOutput
	return output

}
