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

	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range b.OrderChanIn {
		if order.OrderType == "BUY" {
			buyOrders.Push(order)
			if sellOrders.Len() > 0 && sellOrders.Orders[0].Price <= order.Price {
				sellOrder := sellOrders.Pop().(*Order)
				if sellOrder.PendingShare > 0 {
					transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrderChainOut <- sellOrder
					b.OrderChainOut <- order

					if sellOrder.PendingShare > 0 {
						sellOrders.Push(sellOrder)
					}
				}

			}
		} else if order.OrderType == "SELL" {
			sellOrders.Push(order)
		}
	}

}
