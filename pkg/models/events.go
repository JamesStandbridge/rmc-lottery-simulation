package models

import (
	"github.com/ispringteam/eventbus"
)

type TransactionEvent struct {
	From   *Wallet
	To     *Wallet
	Amount float64
}

func (e *TransactionEvent) EventID() eventbus.EventID {
	return "event.transaction"
}

type LotteryOpenEvent struct {
	Lottery *Lottery
}

func (e *LotteryOpenEvent) EventID() eventbus.EventID {
	return "event.lottery.open"
}
