package models

//Create a model for a walle.
//A player has a name and a wallet address
//The wallet also have a balance which is the amount of money the player has

type Wallet struct {
	Address string
	Balance int
}