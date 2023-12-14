package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order         []*Order
	Transactions  []*Transaction
	OrderChanIn   chan *Order
	OrderChainOut chan *Order
	Wg            *sync.WaitGroup
}

func (b *Book) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	sellingShares := transaction.SellingOrder.PendingShare
	buyingShares := transaction.BuyOrder.PendingShare

	minShares := sellingShares
	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.Id, -minShares)
	transaction.SellingOrder.PendingShare -= minShares

	transaction.BuyOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.Id, minShares)
	transaction.BuyOrder.PendingShare -= minShares

	transaction.CalculateTotal(transaction.Shares, transaction.BuyOrder.Price)

	transaction.CloseBuyerOrder()
	transaction.CloseSellingOrder()

	b.Transactions = append(b.Transactions, transaction)
}

func NewBook(orderChanIn, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:         []*Order{},
		Transactions:  []*Transaction{},
		OrderChanIn:   orderChanIn, //receber order do kafka
		OrderChainOut: orderChanOut,
		Wg:            wg,
	}

}

func (b *Book) Trade() {

	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	for order := range b.OrderChanIn {
		asset := order.Asset.Id

		if buyOrders[asset] == nil {
			buyOrders[asset] = NewOrderQueue()
			heap.Init(buyOrders[asset])
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = NewOrderQueue()
			heap.Init(sellOrders[asset])
		}

		if order.OrderType == "BUY" {
			buyOrders[asset].Push(order)
			if sellOrders[asset].Len() > 0 && sellOrders[asset].Orders[0].Price <= order.Price {
				sellOrder := sellOrders[asset].Pop().(*Order)
				if sellOrder.PendingShare > 0 {
					transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrderChainOut <- sellOrder
					b.OrderChainOut <- order

					if sellOrder.PendingShare > 0 {
						sellOrders[asset].Push(sellOrder)
					}
				}

			}
		} else if order.OrderType == "SELL" {
			sellOrders[asset].Push(order)
			if buyOrders[asset].Len() > 0 && buyOrders[asset].Orders[0].Price <= order.Price {
				buyOrder := buyOrders[asset].Pop().(*Order)
				if buyOrder.PendingShare > 0 {
					transaction := NewTransaction(buyOrder, order, order.Shares, buyOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrderChainOut <- buyOrder
					b.OrderChainOut <- order

					if buyOrder.PendingShare > 0 {
						sellOrders[asset].Push(buyOrder)
					}
				}

			}

		}
	}

}
