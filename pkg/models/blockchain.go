package models

import "time"

type Blockchain struct {
	Wallets []*Wallet
	Nfts    []*NFT
}

func (blockchain *Blockchain) Run() []*Wallet {
	for {
		//sleep 1 second
		time.Sleep(1 * time.Second)
		//print "blockchain running"
		println("blockchain running")
	}
}
