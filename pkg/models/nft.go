package models


//Create a new model which represents an NFT
//An NFT has a an array of numbers with a fixed size which represent caracteristics of the NFT
//It also has a owner which is a wallet address

type NFT struct {
	Caracteristics []int
	Owner Wallet
}