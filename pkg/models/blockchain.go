package models

import (
	"math/rand"
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
		chans.TestChan.Publish(&TransactionEvent{
			Amount: 10 + rand.Float64()*(20-10),
		})
	}
}
