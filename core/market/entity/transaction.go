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

func (t *Transaction) CalculateTotal(shares int, price float64) {
	t.Total = float64(shares) * price
}

func (t *Transaction) CloseBuyerOrder() {
	if t.BuyOrder.PendingShare == 0 {
		t.BuyOrder.Status = "CLOSED"
	}
}

func (t *Transaction) CloseSellingOrder() {
	if t.SellingOrder.PendingShare == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}

func NewTransaction(sellingOrder, buyOrder *Order, shares int, price float64) *Transaction {
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyOrder:     buyOrder,
		Shares:       shares,
		Price:        price,
		Total:        float64(shares) * price,
		DateTime:     time.Now(),
	}

}
