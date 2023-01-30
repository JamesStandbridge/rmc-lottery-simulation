package models

import (
	"math/rand"
)

//create a new model for the lottery.
//A lottery is the main class of this application, we can create only one lottery at a time.

type Lottery struct {
	Owner       *Wallet
	Collection  *Collection
	TicketPrice float64
	Tickets     []*NFT
}

// create interface for the lottery
type LotteryInterface interface {
	//the lottery can be initialized
	Initialize(maxDay int)
	//the lottery can be played
	Play()

	Close()
}

func (lottery *Lottery) Close(blockchain Blockchain) {

}

func (lottery *Lottery) Play() {
	currentDay := 0
	maxDay := len(lottery.Collection.Caracteristics)

	for currentDay < maxDay {

		dayCaracteristic := rand.Intn(lottery.Collection.Caracteristics[currentDay])

		for i := 0; i < len(lottery.Tickets); i++ {
			if lottery.Tickets[i].IsActive == false {
				continue
			}
			if currentDay == lottery.Collection.GoldDay {
				lottery.Tickets[i].Type = 1
			}
			if lottery.Tickets[i].Caracteristics[currentDay] != dayCaracteristic {
				lottery.Tickets[i].IsActive = false
			} else if currentDay == maxDay-1 {
				//winner
				lottery.Tickets[i].IsActive = false
				lottery.Tickets[i].Type = 3
			}
		}
		currentDay++
	}
}

// buy a ticket for a wallet
func (lottery *Lottery) BuyTicket(wallet *Wallet, ticketIndex int) {
	//check if index is in range
	if ticketIndex < 0 || ticketIndex >= len(lottery.Tickets) {
		return
	}

	//check if the wallet has enough money
	if wallet.Balance >= lottery.TicketPrice {
		//check if the ticket is still available
		if lottery.Tickets[ticketIndex].Owner.Address == lottery.Owner.Address {
			//transfer the ticket to the wallet
			wallet.Buy(lottery.TicketPrice, lottery.Tickets[ticketIndex])
		}
	}
}
