package dto

type TradeInput struct {
	OrderId      string  `json:"order_id"`
	InvestorId   string  `json:"investor_id"`
	AssetId      string  `json:"asset_id"`
	OrderType    string  `json:"ordertype"`
	CurrentShare int     `json:"current_share"`
	Share        int     `json:"share"`
	Price        float64 `json:"price"`
}

type OrderOutput struct {
	OrderId           string               `json:"order_id"`
	InvestorId        string               `json:"investor_id"`
	AssetId           string               `json:"asset_id"`
	OrderType         string               `json:"ordertype"`
	Status            string               `json:"status"`
	Partial           int                  `json:"partial"`
	Share             int                  `json:"share"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionId string  `json:"transaction_id"`
	BuyerId       string  `json:"buyer_id"`
	SellerId      string  `json:"seller_id"`
	AssetId       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Share         int     `json:"share"`
}
