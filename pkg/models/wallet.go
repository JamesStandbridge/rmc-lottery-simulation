package models

type Wallet struct {
	Address string
	Balance float64
}

// create interface for the lottery
type WalletInterface interface {
	//the lottery can be initialized
	Pay(receiver *Wallet, amount float64)
	//the lottery can be played
	Buy(amount float64, nft *NFT)
}

func (payer *Wallet) Pay(receiver *Wallet, amount float64) {
	if payer.Balance < amount {
		panic("Not enough money")
	} else if amount < 0 {
		panic("Amount must be positive")
	} else if payer == receiver {
		panic("Giver and receiver must be different")
	} else if payer == nil || receiver == nil {
		panic("Giver and receiver must not be nil")
	} else {
		payer.Balance -= amount
		receiver.Balance += amount
	}
}

func (receiver *Wallet) Buy(amount float64, nft *NFT) {
	if receiver.Balance < amount {
		panic("Not enough money")
	} else if amount < 0 {
		panic("Amount must be positive")
	} else if nft == nil {
		panic("NFT must not be nil")
	} else {
		receiver.Pay(nft.Owner, amount)
		nft.Owner = receiver
	}
}
