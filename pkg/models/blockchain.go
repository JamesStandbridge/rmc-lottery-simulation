package models

import (
	"time"
)

type Blockchain struct {
	Wallets []*Wallet
	Nfts    []*NFT
}

func (blockchain *Blockchain) Run(chans BlockchainChanels) {
	//sleep 3 seconds
	for {
		time.Sleep(2 * time.Second)
		chans.LotteryBus.Publish(&LotteryOpenEvent{
			Lottery: &Lottery{},
		})
	}
}
