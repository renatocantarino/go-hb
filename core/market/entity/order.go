package entity

type Order struct {
	ID           string
	Investor     *Investor
	Asset        *Asset
	Shares       int
	PendingShare int
	Price        float64
	OrderType    string
	Status       string
	Transactions []*Transaction
}

func NewOrder(orderId string, investor *Investor, asset *Asset, shares int, pendingShare int,
	price float64, orderType, status string) *Order {

	return &Order{
		ID:           orderId,
		Investor:     investor,
		Asset:        asset,
		Shares:       shares,
		Status:       "OPEN",
		PendingShare: pendingShare,
		OrderType:    orderType,
		Price:        price,
		Transactions: []*Transaction{},
	}

}
