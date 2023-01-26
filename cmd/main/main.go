package main


import (
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/models"
	"github.com/jamesstandbridge/rmc-lottery-simulation/pkg/utils"
)



func main() {
	//create a new wallet
	wallet := models.Wallet{Address: "0x1234567890", Balance: 100}
	//create a new NFT
	
	nfts := GenerateNFTs([]int{2, 2, 2}, wallet)

	//declare an empty array of NFT
	//var nfts []models.NFT

	print(nft.Caracteristics)
}