package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyOrder     *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder, buyOrder *Order, shares int, price float64) *Transaction {
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyOrder:     buyOrder,
		Shares:       shares,
		Price:        price,
		Total:        float64(shares) * price,
	}

}
